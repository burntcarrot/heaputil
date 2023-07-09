package record

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
)

// string: textual description of where this root came from
// uvarint: root pointer
type OtherRootRecord struct {
	Description string
	Address     uint64
}

func (r *OtherRootRecord) GetAddress() uint64 {
	return r.Address
}

func (r *OtherRootRecord) Read(rd *bufio.Reader) error {
	contentSize, err := binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	description := make([]byte, contentSize)
	_, err = io.ReadFull(rd, description)
	if err != nil {
		return err
	}

	r.Description = string(description)

	r.Address, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	return nil
}

func (r *OtherRootRecord) Repr() string {
	format := "OtherRoot at address 0x%x"
	return fmt.Sprintf(format, r.Address)
}
