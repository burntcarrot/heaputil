package record

import (
	"bufio"
	"encoding/binary"
	"fmt"
)

// uvarint: address of this os thread descriptor
// uvarint: Go internal id of thread
// uvarint: os's id for thread
type OSThreadRecord struct {
	Address    uint64
	InternalID uint64
	OSID       uint64
}

func (r *OSThreadRecord) GetAddress() uint64 {
	return r.Address
}

func (r *OSThreadRecord) Read(rd *bufio.Reader) error {
	var err error

	r.Address, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.InternalID, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.OSID, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	return nil
}

func (r *OSThreadRecord) Repr() string {
	format := "OS thread at address 0x%x (Go internal ID = %d, OS ID = %d)"

	return fmt.Sprintf(format, r.Address, r.InternalID, r.OSID)
}
