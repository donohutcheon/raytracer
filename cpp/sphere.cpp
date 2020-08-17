#include "rapidjson/document.h"
#include "sphere.h"

Sphere Sphere::create(const rapidjson::Value::ConstObject &member) {
    auto centerMember = member.FindMember("center");
    auto radiusMember = member.FindMember("radius");
    auto surfaceColorMember = member.FindMember("surfaceColor");
    auto emissionColorMember = member.FindMember("emitColor");
    auto transparencyMember = member.FindMember("transparency");
    auto reflectionMember = member.FindMember("reflection");
    if(centerMember == member.MemberEnd() ||
       radiusMember == member.MemberEnd() ||
       surfaceColorMember == member.MemberEnd() ||
       emissionColorMember == member.MemberEnd() ||
       transparencyMember == member.MemberEnd() ||
       reflectionMember == member.MemberEnd()) {
        throw std::runtime_error("cannot parse bogus sphere");
    }

    Vec3f center = Vec3<double>::create(centerMember->value.GetObject());
    auto radius = radiusMember->value.GetDouble();
    auto surfaceColor = Vec3f::create(surfaceColorMember->value.GetObject());
    auto emissionColor = Vec3f::create(emissionColorMember->value.GetObject());
    auto transparency = transparencyMember->value.GetDouble();
    auto reflection = reflectionMember->value.GetDouble();

    return Sphere(center, radius, surfaceColor, reflection, transparency, emissionColor);
}