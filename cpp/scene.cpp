#include "scene.h"

Scene::Scene(std::shared_ptr<Vec3f> cameraPos, std::shared_ptr<std::vector<Sphere>> spheres)
: cameraPos(move(cameraPos)), spheres(move(spheres))
{}

std::shared_ptr<Scene> Scene::create(const rapidjson::Value::ConstObject &member) {
    auto camera = member.FindMember("camera");
    if(camera == member.MemberEnd()) {
        throw std::runtime_error("cannot parse camera out of bogus scene");
    }
    auto position = camera->value.FindMember("position");
    if(position == camera->value.MemberEnd()) {
        throw std::runtime_error("cannot parse camera position out of bogus scene");
    }
    auto x = position->value.FindMember("x");
    auto y = position->value.FindMember("y");
    auto z = position->value.FindMember("z");
    if(x == position->value.MemberEnd() ||
            y == position->value.MemberEnd() ||
            z == position->value.MemberEnd()) {
        throw std::runtime_error("cannot parse camera co-ordinates out of bogus scene");
    }
    auto cameraPos = std::make_shared<Vec3f>(x->value.GetFloat(), y->value.GetFloat(), z->value.GetFloat());

    auto spheres = member.FindMember("spheres");
    if(spheres == member.MemberEnd()) {
        throw std::runtime_error("cannot parse spheres out of bogus scene");
    }

    auto sphereContainer = std::make_shared<std::vector<Sphere>>() ;
    auto sphereArray = spheres->value.GetArray();
    for(const auto &sphere : sphereArray){
        const rapidjson::Value::ConstObject &s = sphere.GetObject();
        auto s1 = Sphere::create(s);
        sphereContainer->push_back(s1);
    }

    return std::make_shared<Scene>(cameraPos, sphereContainer);
}