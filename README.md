# heaputil

Library for parsing Go heap dumps.

Check [example](./example/main.go).

Output:

```
Created a sample heapdump.
DumpParams (big endian = false, ptrsize = 8, heap start = 0xc000000000, heap end = 0xc004000000, architecture = amd64, GOEXPERIMENT = go1.20, numCPU = 5)
TypeDescriptor at address 0x4a1900 (name = github.com/burntcarrot/heaputil/record., size = 8, hasType = false)
ITab at address 0x4d2508, type descriptor address = 0x4a1900
TypeDescriptor at address 0x4a3180 (name = github.com/burntcarrot/heaputil/record., size = 8, hasType = false)
ITab at address 0x4d27b0, type descriptor address = 0x4a3180
TypeDescriptor at address 0x4a4ee0 (name = io.discard, size = 0, hasType = false)
ITab at address 0x4d2488, type descriptor address = 0x4a4ee0
TypeDescriptor at address 0x4abd80 (name = encoding/binary.littleEndian, size = 0, hasType = false)
ITab at address 0x4d2bb8, type descriptor address = 0x4abd80
```
