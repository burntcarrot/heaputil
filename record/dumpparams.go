package record

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
)

// bool: big endian
// uvarint: pointer size in bytes
// uvarint: starting address of heap
// uvarint: ending address of heap
// string: architecture name
// string: GOEXPERIMENT environment variable value
// uvarint: runtime.ncpu
type DumpParamsRecord struct {
	BigEndian     bool
	PtrSize       uint64
	HeapStartAddr uint64
	HeapEndAddr   uint64
	Architecture  string
	GoExperiment  string
	NumCPU        uint64
}

func (r *DumpParamsRecord) Read(rd *bufio.Reader) error {
	bigEndian, err := binary.ReadUvarint(rd)
	if err != nil {
		return err
	}
	r.BigEndian = (bigEndian != 0)

	r.PtrSize, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.HeapStartAddr, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.HeapEndAddr, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	contentSize, err := binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	architecture := make([]byte, contentSize)
	_, err = io.ReadFull(rd, architecture)
	if err != nil {
		return err
	}

	r.Architecture = string(architecture)

	contentSize, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	goExperiment := make([]byte, contentSize)
	_, err = io.ReadFull(rd, goExperiment)
	if err != nil {
		return err
	}

	r.GoExperiment = string(goExperiment)

	r.NumCPU, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	return nil
}

func (r *DumpParamsRecord) Repr() string {
	format := "DumpParams (big endian = %v, ptrsize = %d, heap start = 0x%x, heap end = 0x%x, architecture = %s, GOEXPERIMENT = %s, numCPU = %d)"
	return fmt.Sprintf(format, r.BigEndian, r.PtrSize, r.HeapStartAddr, r.HeapEndAddr, r.Architecture, r.GoExperiment, r.NumCPU)
}
