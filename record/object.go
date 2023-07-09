package record

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
)

// uvarint: address of object
// string: contents of object
// fieldlist: describes pointer-containing fields of the object
type ObjectRecord struct {
	Address  uint64
	Contents []byte
	Fields   []uint64
}

func (o *ObjectRecord) GetAddress() uint64 {
	return o.Address
}

func (o *ObjectRecord) GetContent() []byte {
	return o.Contents
}

func (o *ObjectRecord) GetFields() []uint64 {
	return o.Fields
}

func (o *ObjectRecord) Repr() string {
	format := "Object at address 0x%x (pointers=%d) (bytes=%d)"
	return fmt.Sprintf(format, o.Address, len(o.Fields), len(o.Contents))
}

func (o *ObjectRecord) Read(rd *bufio.Reader) error {
	var err error
	o.Address, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	contentSize, err := binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	o.Contents = make([]byte, contentSize)
	_, err = io.ReadFull(rd, o.Contents)
	if err != nil {
		return err
	}

	o.Fields, err = ReadFieldList(rd)
	if err != nil {
		return err
	}

	return nil
}
