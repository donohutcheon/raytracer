#include "image.h"

std::shared_ptr<Image> Image::create(const rapidjson::Value::ConstObject &member) {
    auto width = member.FindMember("width");
    if(width == member.MemberEnd()) {
        throw std::runtime_error("cannot parse bogus image; could not find image width");
    }

    auto height = member.FindMember("height");
    if(height == member.MemberEnd()) {
        throw std::runtime_error("cannot parse bogus image; could not find image height");
    }
    int w = width->value.GetInt();
    int h = height->value.GetInt();

    return std::make_shared<Image>(w, h);
}