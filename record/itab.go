package record

import (
	"bufio"
	"encoding/binary"
	"fmt"
)

// uvarint: Itab address
// uvarint: address of type descriptor for contained type
type ITabRecord struct {
	Address            uint64
	TypeDescriptorAddr uint64
}

func (r *ITabRecord) GetAddress() uint64 {
	return r.Address
}

func (r *ITabRecord) Read(rd *bufio.Reader) error {
	var err error

	r.Address, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.TypeDescriptorAddr, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	return nil
}

func (r *ITabRecord) Repr() string {
	format := "ITab at address 0x%x, type descriptor address = 0x%x"

	return fmt.Sprintf(format, r.Address, r.TypeDescriptorAddr)
}
