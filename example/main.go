package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/debug"

	"github.com/burntcarrot/heaputil"
)

func main() {
	// Simulating some memory allocations
	for i := 0; i < 10; i++ {
		_ = make([]byte, 1024)
	}

	f, err := os.Create("heapdump")
	if err != nil {
		panic("Could not open file for writing:" + err.Error())
	}

	// Trigger GC to clean up unreachable objects
	// not required, but recommended before taking a heap dump
	runtime.GC()

	// Perform a heap dump.
	debug.WriteHeapDump(f.Fd())
	f.Close()

	fmt.Println("Created a sample heapdump.")

	file, err := os.Open("heapdump")
	if err != nil {
		log.Fatalln("failed to open heap dump")
	}

	reader := bufio.NewReader(file)

	// Use heaputil's PrintDump to print the dump.
	// CAUTION: This will print a lot of information to stdout!
	err = heaputil.PrintDump(reader)
	if err != nil {
		log.Fatalf("failed to print heap dump: %v\n", err)
	}
}
