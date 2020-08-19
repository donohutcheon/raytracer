#ifndef RAYTRACER_SCENE_H
#define RAYTRACER_SCENE_H
#include <memory>
#include <vector>

#include "image.h"
#include "sphere.h"

class Scene {
    std::shared_ptr<Vec3f> cameraPos;
    std::shared_ptr<std::vector<Sphere>> spheres;
public:
    Scene(std::shared_ptr<Vec3f> cameraPos, std::shared_ptr<std::vector<Sphere>> spheres);
    static std::shared_ptr<Scene> create(const rapidjson::Value::ConstObject &member);
    std::shared_ptr<Vec3f> getCameraPos(){return cameraPos;}
    std::shared_ptr<std::vector<Sphere>> getSpheres(){return spheres;}
};

#endif //RAYTRACER_SCENE_H
