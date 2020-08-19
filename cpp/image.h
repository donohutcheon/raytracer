#ifndef RAYTRACER_IMAGE_H
#define RAYTRACER_IMAGE_H

#include <memory>
#include "rapidjson/document.h"

class Image {
    int width;
    int height;
public:
    Image(int width, int height) : width(width), height(height){}
    int getWidth() { return width;}
    int getHeight() { return height;}
    static std::shared_ptr<Image> create(const rapidjson::Value::ConstObject &member);
};


#endif //RAYTRACER_IMAGE_H
