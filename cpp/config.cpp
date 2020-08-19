#include <memory>
#include "config.h"



std::shared_ptr<Config> Config::create(const rapidjson::Document &doc) {

    auto image = doc.FindMember("image");
    if(image == doc.MemberEnd()) {
        throw std::runtime_error("cannot parse image out of bogus config");
    }
    auto scene = doc.FindMember("scene");
    if(scene == doc.MemberEnd()) {
        throw std::runtime_error("cannot parse scene out of bogus config");
    }

    auto i = Image::create(image->value.GetObject());
    auto s = Scene::create(scene->value.GetObject());

    return std::make_shared<Config>(i, s);
}