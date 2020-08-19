#include <cstdlib>
#include <cstdio>
#include <cmath>
#include <fstream>
#include <memory>
#include <vector>
#include <iostream>
#include <sstream>

#include "rapidjson/document.h"
#include "rapidjson/istreamwrapper.h"
#include "rapidjson/stringbuffer.h"

#include "sphere.h"
#include "config.h"


#if defined __linux__ || defined __APPLE__
// "Compiled for Linux
#else
// Windows doesn't define these values by default, Linux does
#define M_PI 3.141592653589793
#define INFINITY 1e8
#endif

#define MAX_RAY_DEPTH 5

double mix(const double &a, const double &b, const double &mix)
{
    return b * mix + a * (1 - mix);
}
Vec3f trace(
        const Vec3f &rayorig,
        const Vec3f &raydir,
        const std::vector<Sphere> &spheres,
        const int &depth)
{
    //if (raydir.length() != 1) std::cerr << "Error " << raydir << std::endl;
    double tnear = INFINITY;
    const Sphere* sphere = NULL;
    // find intersection of this ray with the sphere in the scene
    for (unsigned i = 0; i < spheres.size(); ++i) {
        double t0 = INFINITY, t1 = INFINITY;
        if (spheres[i].intersect(rayorig, raydir, t0, t1)) {
            if (t0 < 0) t0 = t1;
            if (t0 < tnear) {
                tnear = t0;
                sphere = &spheres[i];
            }
        }
    }
    // if there's no intersection return black or background color
    if (!sphere)
        return Vec3f(2);
    Vec3f surfaceColor = 0; // color of the ray/surfaceof the object intersected by the ray
    Vec3f phit = rayorig + raydir * tnear; // point of intersection
    Vec3f nhit = phit - sphere->center; // normal at the intersection point
    nhit.normalize(); // normalize normal direction
    // If the normal and the view direction are not opposite to each other
    // reverse the normal direction. That also means we are inside the sphere so set
    // the inside bool to true. Finally reverse the sign of IdotN which we want
    // positive.
    double bias = 1e-4; // add some bias to the point from which we will be tracing
    bool inside = false;
    if (raydir.dot(nhit) > 0) nhit = -nhit, inside = true;
    if ((sphere->transparency > 0 || sphere->reflection > 0) && depth < MAX_RAY_DEPTH) {
        double facingratio = -raydir.dot(nhit);
        // change the mix value to tweak the effect
        double fresneleffect = mix(pow(1 - facingratio, 3), 1, 0.1);
        // compute reflection direction (not need to normalize because all vectors
        // are already normalized)
        Vec3f refldir = raydir - nhit * 2 * raydir.dot(nhit);
        refldir.normalize();
        Vec3f orig = phit + nhit * bias;
        Vec3f reflection = trace(orig, refldir, spheres, depth + 1);
        Vec3f refraction = 0;
        // if the sphere is also transparent compute refraction ray (transmission)
        if (sphere->transparency) {
            double ior = 1.1, eta = (inside) ? ior : 1 / ior; // are we inside or outside the surface?
            double cosi = -nhit.dot(raydir);
            double k = 1 - eta * eta * (1 - cosi * cosi);
            Vec3f refrdir = raydir * eta + nhit * (eta *  cosi - sqrt(k));
            refrdir.normalize();
            Vec3f newOrigin = phit - nhit * bias;
            refraction = trace(newOrigin, refrdir, spheres, depth + 1);
        }
        // the result is a mix of reflection and refraction (if the sphere is transparent)
        surfaceColor = (
                               reflection * fresneleffect +
                               refraction * (1 - fresneleffect) * sphere->transparency) * sphere->surfaceColor;
    }
    else {
        // it's a diffuse object, no need to raytrace any further
        for (unsigned i = 0; i < spheres.size(); ++i) {
            if (spheres[i].emissionColor.x > 0) {
                // this is a light
                Vec3f transmission = 1;
                Vec3f lightDirection = spheres[i].center - phit;
                lightDirection.normalize();
                for (unsigned j = 0; j < spheres.size(); ++j) {
                    if (i != j) {
                        double t0, t1;
                        if (spheres[j].intersect(phit + nhit * bias, lightDirection, t0, t1)) {
                            transmission = 0;
                            break;
                        }
                    }
                }
                double max = std::max(double(0), nhit.dot(lightDirection));
                surfaceColor += sphere->surfaceColor * transmission *
                                max * spheres[i].emissionColor;
            }
        }
    }

    return surfaceColor + sphere->emissionColor;
}

Vec3f castRay(int x, int y, const std::vector<Sphere> &spheres) {
    unsigned width = 640, height = 480;
    double invWidth = 1 / double(width), invHeight = 1 / double(height);
    double fov = 30, aspectratio = width / double(height);
    double angle = tan(M_PI * 0.5 * fov / 180.);
    double xx = (2 * ((x + 0.5) * invWidth) - 1) * angle * aspectratio;
    double yy = (1 - 2 * ((y + 0.5) * invHeight)) * angle;
    Vec3f raydir(xx, yy, -1);
    raydir.normalize();
    Vec3f pixel = trace(Vec3f(0), raydir, spheres, 0);

    return pixel;
}

void render(std::shared_ptr<Config> config)
{
    auto spheres = config->getScene()->getSpheres();

    unsigned width = config->getImage()->getWidth();
    unsigned height = config->getImage()->getHeight();
    auto cameraPos = config->getScene()->getCameraPos();

    Vec3f *image = new Vec3f[width * height], *pixel = image;
    double invWidth = 1 / double(width), invHeight = 1 / double(height);
    double fov = 30, aspectratio = width / double(height);
    double angle = tan(M_PI * 0.5 * fov / 180.);
    // Trace rays
    for (unsigned y = 0; y < height; ++y) {
        for (unsigned x = 0; x < width; ++x, ++pixel) {
            double xx = (2 * ((x + 0.5) * invWidth) - 1) * angle * aspectratio;
            // ((2 * x * invWidth) + (1 * invWidth)) - 1) * angle * aspectratio;
            double yy = (1 - 2 * ((y + 0.5) * invHeight)) * angle;
            Vec3f raydir(xx, yy, -1);
            //std::cout << x << " " << y << " " << raydir << std::endl;
            raydir.normalize();
            *pixel = trace(*cameraPos, raydir, *spheres, 0);
        }
    }
    // Save result to a PPM image (keep these flags if you compile under Windows)
    std::ofstream ofs("./untitled.ppm", std::ios::out | std::ios::binary);
    ofs << "P6\n" << width << " " << height << "\n255\n";
    for (unsigned i = 0; i < width * height; ++i) {
        ofs << (unsigned char)(std::min(double(1), image[i].x) * 255) <<
            (unsigned char)(std::min(double(1), image[i].y) * 255) <<
            (unsigned char)(std::min(double(1), image[i].z) * 255);
    }
    ofs.close();
    delete [] image;
}

std::shared_ptr<Config> loadScene(std::string filename) {

    std::ifstream ifs { filename };
    if ( !ifs.is_open() )
    {
        throw std::runtime_error("could not open file for reading");
    }

    rapidjson::IStreamWrapper isw { ifs };
    rapidjson::Document doc {};
    doc.ParseStream( isw );
    ifs.close();

    rapidjson::StringBuffer buffer {};

    if ( doc.HasParseError() )
    {
        std::stringstream error;
        error << "Error  : " << doc.GetParseError()  << '\n'
                  << "Offset : " << doc.GetErrorOffset() << '\n';
        throw std::runtime_error(error.str());
    }

    const std::string jsonStr { buffer.GetString() };

    return Config::create(doc);
}

int main(int argc, char **argv)
{
    std::string path;
    if(argc != 2) {
        return EXIT_FAILURE;
    }
    std::string sceneFile = argv[1];
    std::cout << sceneFile << std::endl;

    try {
        auto config = loadScene(sceneFile);
        srand48(13);
        render(config);
    } catch (std::exception &e) {
        std::cerr << e.what() << std::endl;
        return EXIT_FAILURE;
    }

    return 0;
}
