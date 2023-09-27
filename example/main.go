package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime/debug"

	"github.com/burntcarrot/heaputil"
)

func main() {
	for i := 0; i < 10; i++ {
		_ = make([]byte, 1024)
	}

	f, err := os.Create("heapdump")
	if err != nil {
		log.Fatalf("Could not open file for writing: %v\n", err)
	}

	debug.WriteHeapDump(f.Fd())
	f.Close()

	fmt.Println("Created a sample heapdump.")

	file, err := os.Open("heapdump")
	if err != nil {
		log.Fatalln("failed to open heap dump")
	}

	reader := bufio.NewReader(file)

	// Use heaputil's ParseDump to get records
	records, err := heaputil.ParseDump(reader)
	if err != nil {
		log.Fatalf("failed to parse heap dump: %v\n", err)
	}

	// Print each record.
	for _, r := range records {
		fmt.Printf("%s\n", r.Repr)
	}
}
