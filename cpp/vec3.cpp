#include <cmath>
#include "vec3.h"

template<typename T>
Vec3<T>& Vec3<T>::normalize()
{
    T nor2 = length2();
    if (nor2 > 0) {
        T invNor = 1 / sqrt(nor2);
        x *= invNor, y *= invNor, z *= invNor;
    }
    return *this;
}


template<typename T>
Vec3<double> Vec3<T>::create(const rapidjson::Value::ConstObject &member) {
    auto xMember = member.FindMember("x");
    auto yMember = member.FindMember("y");
    auto zMember = member.FindMember("z");

    if(xMember == member.MemberEnd() ||
       yMember == member.MemberEnd() ||
       zMember == member.MemberEnd()) {
        throw std::runtime_error("cannot parse bogus vector3");
    }
    if(!xMember->value.IsNumber() ||
       !yMember->value.IsNumber() ||
       !zMember->value.IsNumber()) {
        throw std::runtime_error("cannot parse bogus vector3 with non-numeric values");
    }
    auto x = xMember->value.GetDouble();
    auto y = yMember->value.GetDouble();
    auto z = zMember->value.GetDouble();

    return Vec3<double>(x, y, z);
}

template class Vec3<double>;