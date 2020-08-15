# Raytracer


Inline-style: 
![alt text](https://github.com/donohutcheon/raytracer/blob/master/docs/example.png "Example scene")

A simple raytracer implemented in Go to explore the profiling and optimization tooling that is provided by Go.

## Original source

The source was apdapted from C++ source obtained from www.scratchapixel.com; released under GPL V3 or later.

## Running and Analysing

The following command will build and run the raytracer.
```
go build && time ./raytracer
```

Use the following command to build, run and then analyse the cpu.pprof profile file.
```
go build && time ./raytracer &&  go tool pprof -http=:8080 cpu.pprof
```

## Useful resources and links

[GopherCon 2019: Dave Cheney - Two Go Programs, Three Different Profiling Techniques](https://youtu.be/nok0aYiGiYA)

[#22: using the Go execution tracer](https://www.youtube.com/watch?v=ySy3sR1LFCQ) 