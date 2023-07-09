package record

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
)

// uvarint: address of type descriptor
// uvarint: size of an object of this type
// string: name of type
// bool: whether the data field of an interface containing a value of this type has type T (false) or *T (true)
type TypeDescriptorRecord struct {
	Address uint64
	Size    uint64
	Name    string
	HasType bool
}

func (r *TypeDescriptorRecord) GetAddress() uint64 {
	return r.Address
}

func (r *TypeDescriptorRecord) Read(rd *bufio.Reader) error {
	var err error
	r.Address, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.Size, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	contentSize, err := binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	name := make([]byte, contentSize)
	_, err = io.ReadFull(rd, name)
	if err != nil {
		return err
	}

	r.Name = string(name)

	hasTypeInt, err := binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.HasType = (hasTypeInt != 0)

	return nil
}

func (r *TypeDescriptorRecord) Repr() string {
	format := "TypeDescriptor at address 0x%x (name = %s, size = %d, hasType = %v)"
	return fmt.Sprintf(format, r.Address, r.Name, r.Size, r.HasType)
}
