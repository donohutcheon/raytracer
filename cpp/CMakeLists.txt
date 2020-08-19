cmake_minimum_required(VERSION 3.17)
project(raytracer)

set(CMAKE_CXX_STANDARD 20)

include(${CMAKE_BINARY_DIR}/conanbuildinfo.cmake)
conan_basic_setup()

add_executable(raytracer raytracer.cpp vec3.h vec3.cpp sphere.cpp sphere.h scene.cpp scene.h image.cpp image.h config.cpp config.h)
target_link_libraries(raytracer ${CONAN_LIBS})