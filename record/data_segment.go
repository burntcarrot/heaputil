package record

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
)

// uvarint: address of the start of the data segment
// string: contents of the data segment
// fieldlist: kind and offset of pointer-containing fields in the data segment.
type DataSegmentRecord struct {
	Address  uint64
	Contents []byte
	Fields   []uint64
}

func (r *DataSegmentRecord) GetAddress() uint64 {
	return r.Address
}

func (r *DataSegmentRecord) GetContent() []byte {
	return r.Contents
}

func (r *DataSegmentRecord) GetFields() []uint64 {
	return r.Fields
}

func (r *DataSegmentRecord) Read(rd *bufio.Reader) error {
	var err error

	r.Address, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	contentSize, err := binary.ReadUvarint(rd)
	if err != nil {
		return err
	}
	contents := make([]byte, contentSize)
	_, err = io.ReadFull(rd, contents)
	if err != nil {
		return err
	}
	r.Contents = contents

	r.Fields, err = ReadFieldList(rd)
	if err != nil {
		return err
	}

	return nil
}

func (r *DataSegmentRecord) Repr() string {
	format := "Data segment at address 0x%x (content size = %d, pointers = %d)"

	return fmt.Sprintf(format, r.Address, len(r.Contents), len(r.Fields))
}
