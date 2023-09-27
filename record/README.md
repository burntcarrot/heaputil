# Records

The heap dump has the header `go1.7 heap dump\n`, and the rest of the dump is composed of records. Each record starts with a 64-bit unsigned integer describing the type of the record:

- 0 = EOF
- 1 = object
- 2 = otherroot
- 3 = type
- 4 = goroutine
- 5 = stack frame
- 6 = dump params
- 7 = registered finalizer
- 8 = itab
- 9 = OS thread
- 10 = mem stats
- 11 = queued finalizer
- 12 = data segment
- 13 = bss segment
- 14 = defer record
- 15 = panic record
- 16 = alloc/free profile record
- 17 = alloc stack trace sample

(Explained in detail here: [https://github.com/golang/go/wiki/heapdump15-through-heapdump17](https://github.com/golang/go/wiki/heapdump15-through-heapdump17))

In this package, all of these records are defined as structs. For example:

```go
type ObjectRecord struct {
    Address uint64
    // ...
}
```

and these records implement the `Record` interface:

```go
type Record interface {
	Read(r *bufio.Reader) error
	Repr() string
}
```

The idea is to use a single reader, and pass it on several records, who do their processing inside `Read`. All records need to have a representation through `Repr()`.
