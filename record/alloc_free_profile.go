package record

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
)

// uvarint: record identifier
// uvarint: size of allocated object
// uvarint: number of stack frames.
// uvarint: number of allocations
// uvarint: number of frees
type AllocFreeProfileRecord struct {
	ID         uint64
	Size       uint64
	FrameCount uint64
	AllocCount uint64
	FreeCount  uint64
	Frames     []Frame
}

// Each frame contains:
// string: function name
// string: file name
// uvarint: line number
type Frame struct {
	Name     string
	Filename string
	Line     uint64
}

func (r *AllocFreeProfileRecord) Read(rd *bufio.Reader) error {
	var err error

	r.ID, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.Size, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.FrameCount, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.Frames = make([]Frame, r.FrameCount)
	for i := uint64(0); i < r.FrameCount; i++ {
		contentSize, err := binary.ReadUvarint(rd)
		if err != nil {
			return err
		}

		name := make([]byte, contentSize)
		_, err = io.ReadFull(rd, name)
		if err != nil {
			return err
		}

		r.Frames[i].Name = string(name)

		contentSize, err = binary.ReadUvarint(rd)
		if err != nil {
			return err
		}

		filename := make([]byte, contentSize)
		_, err = io.ReadFull(rd, filename)
		if err != nil {
			return err
		}

		r.Frames[i].Filename = string(filename)

		r.Frames[i].Line, err = binary.ReadUvarint(rd)
		if err != nil {
			return err
		}
	}

	r.AllocCount, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.FreeCount, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	return nil
}

func (r *AllocFreeProfileRecord) Repr() string {
	format := "Alloc/free profile record: %+v"

	return fmt.Sprintf(format, *r)
}
