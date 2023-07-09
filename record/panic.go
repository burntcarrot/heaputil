package record

import (
	"bufio"
	"encoding/binary"
	"fmt"
)

// uvarint: panic record address
// uvarint: containing goroutine
// uvarint: type ptr of panic arg eface
// uvarint: data field of panic arg eface
// uvarint: ptr to defer record that's currently running
// uvarint: link to next panic record
type PanicRecordRecord struct {
	Address           uint64
	ContainsGoroutine uint64
	ArgType           uint64
	ArgData           uint64
	DeferPtr          uint64
	Next              uint64
}

func (r *PanicRecordRecord) GetAddress() uint64 {
	return r.Address
}

func (r *PanicRecordRecord) Read(rd *bufio.Reader) error {
	var err error

	r.Address, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.ContainsGoroutine, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.ArgType, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.ArgData, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.DeferPtr, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.Next, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	return nil
}

func (r *PanicRecordRecord) Repr() string {
	format := "Panic record at address 0x%x (contains goroutine = %d, arg type = %d, defer record ptr = %d, next defer record = %d)"

	return fmt.Sprintf(format, r.Address, r.ContainsGoroutine, r.ArgType, r.DeferPtr, r.Next)
}
