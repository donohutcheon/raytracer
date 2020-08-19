#ifndef RAYTRACER_CONFIG_H
#define RAYTRACER_CONFIG_H

#include <memory>
#include "image.h"
#include "scene.h"

class Config {
    std::shared_ptr<Image> image;
    std::shared_ptr<Scene> scene;
public:
    Config(std::shared_ptr<Image> image, std::shared_ptr<Scene> scene)
    : image(move(image)), scene(move(scene)) {}
    static std::shared_ptr<Config> create(const rapidjson::Document &doc);
    std::shared_ptr<Image> getImage() {return image;}
    std::shared_ptr<Scene> getScene() {return scene;}
};


#endif //RAYTRACER_CONFIG_H
