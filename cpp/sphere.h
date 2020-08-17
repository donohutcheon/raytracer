#ifndef RAYTRACER_SPHERE_H
#define RAYTRACER_SPHERE_H

#include <cmath>
#include "rapidjson/document.h"
#include "vec3.h"

class Sphere
{
public:
    Vec3f center;                           /// position of the sphere
    double radius, radius2;                  /// sphere radius and radius^2
    Vec3f surfaceColor, emissionColor;      /// surface color and emission (light)
    double transparency, reflection;         /// surface transparency and reflectivity
    Sphere(
            const Vec3f &c,
            const double &r,
            const Vec3f &sc,
            const double &refl = 0,
            const double &transp = 0,
            const Vec3f &ec = 0) :
            center(c), radius(r), radius2(r * r), surfaceColor(sc), emissionColor(ec),
            transparency(transp), reflection(refl)
    { /* empty */ }
    static Sphere create(const rapidjson::Value::ConstObject &member);
    bool intersect(const Vec3f &rayorig, const Vec3f &raydir, double &t0, double &t1) const
    {
        Vec3f l = center - rayorig;
        double tca = l.dot(raydir);
        if (tca < 0)
            return false;
        double d2 = l.dot(l) - tca * tca;
        if (d2 > radius2) return false;
        double thc = sqrt(radius2 - d2);
        t0 = tca - thc;
        t1 = tca + thc;

        return true;
    }
    friend std::ostream & operator << (std::ostream &os, const Sphere &s)
    {
        os << "center: " << s.center << ", ";
        os << "radius: " << s.radius << ", ";
        os << "surfaceColor: " << s.surfaceColor << ", ";
        os << "transparency: " << s.transparency << ", ";
        os << "reflection: " << s.reflection << ", ";
        os << "emissionColor: " << s.emissionColor;
        return os;
    }
};

#endif