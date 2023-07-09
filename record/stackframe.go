package record

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
)

// uvarint: stack pointer (lowest address in frame)
// uvarint: depth in stack (0 = top of stack)
// uvarint: stack pointer of child frame (or 0 if none)
// string: contents of stack frame
// uvarint: entry pc for function
// uvarint: current pc for function
// uvarint: continuation pc for function (where function may resume, if anywhere)
// string: function name
// fieldlist: list of kind and offset of pointer-containing fields in this frame
type StackFrameRecord struct {
	Address        uint64
	Depth          uint64
	ChildPtr       uint64
	Contents       []byte
	EntryPC        uint64
	CurrentPC      uint64
	ContinuationPC uint64
	Name           string
	Fields         []uint64
}

func (r *StackFrameRecord) GetAddress() uint64 {
	return r.Address
}

func (r *StackFrameRecord) GetFields() []uint64 {
	return r.Fields
}

func (r *StackFrameRecord) GetContent() []byte {
	return r.Contents
}

func (r *StackFrameRecord) Read(rd *bufio.Reader) error {
	var err error
	r.Address, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.Depth, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.ChildPtr, err = binary.ReadUvarint(rd)
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

	r.EntryPC, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.CurrentPC, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.ContinuationPC, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	contentSize, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	name := make([]byte, contentSize)
	_, err = io.ReadFull(rd, name)
	if err != nil {
		return err
	}

	r.Name = string(name)

	r.Fields, err = ReadFieldList(rd)
	if err != nil {
		return err
	}

	return nil
}

func (r *StackFrameRecord) Repr() string {
	format := "Stack Frame [%s] at address 0x%x (depth = %d, pointers = %d, content bytes = %d, child pointer address = 0x%x)"

	return fmt.Sprintf(format, r.Name, r.Address, r.Depth, len(r.Fields), len(r.Contents), r.ChildPtr)
}
