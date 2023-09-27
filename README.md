# heaputil

Library for parsing Go heap dumps.

Check [example](./example/main.go).

Output:

```
....
OS thread at address 0x54b300 (Go internal ID = 0, OS ID = 53413)
Data segment at address 0x546660 (content size = 16784, pointers = 1184)
BSS segment at address 0x54a800 (content size = 196376, pointers = 10185)
Registered finalizer at address 0xc000078060, FuncVal ptr address = 0x4b7078, type = 4832448, object type = 4832448
Registered finalizer at address 0xc0000a8000, FuncVal ptr address = 0x4b7078, type = 4832448, object type = 4832448
Registered finalizer at address 0xc0000a8060, FuncVal ptr address = 0x4b7078, type = 4832448, object type = 4832448
Registered finalizer at address 0xc0000a80c0, FuncVal ptr address = 0x4b7078, type = 4832448, object type = 4832448
....
```
