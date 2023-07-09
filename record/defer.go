package record

import (
	"bufio"
	"encoding/binary"
	"fmt"
)

// uvarint: defer record address
// uvarint: containing goroutine
// uvarint: argp
// uvarint: pc
// uvarint: FuncVal of defer
// uvarint: PC of defer entry point
// uvarint: link to next defer record
type DeferRecordRecord struct {
	Address           uint64
	ContainsGoroutine uint64
	Argp              uint64
	PC                uint64
	FuncVal           uint64
	EntrypointPC      uint64
	Next              uint64
}

func (r *DeferRecordRecord) GetAddress() uint64 {
	return r.Address
}

func (r *DeferRecordRecord) Read(rd *bufio.Reader) error {
	var err error

	r.Address, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.ContainsGoroutine, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.Argp, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.PC, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.FuncVal, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.EntrypointPC, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.Next, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	return nil
}

func (r *DeferRecordRecord) Repr() string {
	format := "Defer record at address 0x%x (contains goroutine = %d, argp = %d, FuncVal = %d, next defer record = %d)"

	return fmt.Sprintf(format, r.Address, r.ContainsGoroutine, r.Argp, r.FuncVal, r.Next)
}
