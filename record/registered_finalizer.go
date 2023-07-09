package record

import (
	"bufio"
	"encoding/binary"
	"fmt"
)

// uvarint: address of object that has a finalizer
// uvarint: pointer to FuncVal describing the finalizer
// uvarint: PC of finalizer entry point
// uvarint: type of finalizer argument
// uvarint: type of object
type RegisteredFinalizerRecord struct {
	ObjectAddress uint64
	Address       uint64
	EntryPC       uint64
	Type          uint64
	ObjectType    uint64
}

func (r *RegisteredFinalizerRecord) Read(rd *bufio.Reader) error {
	var err error

	r.ObjectAddress, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.Address, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.EntryPC, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.Type, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.ObjectType, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	return nil
}

func (r *RegisteredFinalizerRecord) Repr() string {
	format := "Registered finalizer at address 0x%x, FuncVal ptr address = 0x%x, type = %d, object type = %d"

	return fmt.Sprintf(format, r.ObjectAddress, r.Address, r.Type, r.ObjectType)
}
