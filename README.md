# Raytracer
 
![alt text](https://github.com/donohutcheon/raytracer/blob/master/docs/example.png "Example scene")

A simple raytracer implemented in Go to explore the profiling and optimization tooling that is provided by Go.

## Original source

The source was adapted from C++ source obtained from www.scratchapixel.com; released under GPL V3 or later.

## Running and Analysing

The following command will build and run the raytracer.
```
go build && time ./raytracer
```

Use the following command to build, run and then analyse the cpu.pprof profile file.
```
go build && time ./raytracer &&  go tool pprof -http=:8080 cpu.pprof
```

Benchmark tests can be utilized to collect profiles.  Go test runs in two different modes:
`local directory mode` and `package list mode`.  It seems that benchmark profiling only 
works in `local directory mode`; thus you need to change directory into the package you wish to 
profile. Refer to [Package test](https://golang.org/pkg/cmd/go/internal/test/)

Use the following command to run a benchmark test while profiling the CPU.
```
cd raytrace/
go test -run=. -bench=. -cpuprofile=cpu.out
cd -
go tool pprof -http=: raytrace/cpu.out
```

Use the following command to run a benchmark test while profiling memory allocations.
The `benchmem` argument prints additional memory allocation statistics for benchmarks.
```
cd raytrace/
go test -run=. -bench=. -benchmem -memprofile=mem.out
cd -
go tool pprof -http=: raytrace/mem.out
```

Use the following command to run a benchmark test while running the trace profiler.
```
cd raytrace/
go test -run=. -bench=. -trace=trace.out
cd -
go tool trace raytrace/trace.out
```

It's possible to execute multiple profile types per test. 
```
cd raytrace/
go test -run=. -bench=. -benchmem -cpuprofile=cpu.out -memprofile=mem.out -trace=trace.out
cd -
```

## Building and running the C++ Raytracer

Conan package manager and CMake are required to compile and build the program.
Use the following commands to build and run the C++ version of the raytracer.
```
cd cpp
conan install .
cmake . -G "Unix Makefiles" -DCMAKE_BUILD_TYPE=Release
cmake --build .
time bin/raytracer ../scene.json 
```

## Useful resources and links

[GopherCon 2019: Dave Cheney - Two Go Programs, Three Different Profiling Techniques](https://youtu.be/nok0aYiGiYA)
[#22: using the Go execution tracer](https://www.youtube.com/watch?v=ySy3sR1LFCQ) 