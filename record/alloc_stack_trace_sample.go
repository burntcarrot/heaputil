package record

import (
	"bufio"
	"encoding/binary"
	"fmt"
)

// uvarint: address of object
// uvarint: alloc/free profile record identifier
type AllocStackTraceSampleRecord struct {
	Address  uint64
	RecordID uint64
}

func (r *AllocStackTraceSampleRecord) GetAddress() uint64 {
	return r.Address
}

func (r *AllocStackTraceSampleRecord) Read(rd *bufio.Reader) error {
	var err error

	r.Address, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.RecordID, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	return nil
}

func (r *AllocStackTraceSampleRecord) Repr() string {
	format := "Alloc stack trace sample at address 0x%x, record ID = %d"

	return fmt.Sprintf(format, r.Address, r.RecordID)
}
