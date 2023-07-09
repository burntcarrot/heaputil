package record

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
)

// uvarint: address of descriptor
// uvarint: pointer to the top of stack (the currently running frame, a.k.a. depth 0)
// uvarint: go routine ID
// uvarint: the location of the go statement that created this goroutine
// uvarint: status
// bool: is a Go routine started by the system
// bool: is a background Go routine
// uvarint: approximate time the go routine last started waiting (nanoseconds since the Epoch)
// string: textual reason why it is waiting
// uvarint: context pointer of currently running frame
// uvarint: address of os thread descriptor (M)
// uvarint: top defer record
// uvarint: top panic record
type GoroutineRecord struct {
	Address       uint64
	StackPtr      uint64
	RoutineID     uint64
	CreatorPtr    uint64
	Status        Status
	System        bool
	Background    bool
	WaitTime      uint64
	WaitReason    string
	ContextPtr    uint64
	ThreadAddress uint64
	TopDefer      uint64
	TopPanic      uint64
}

func (r *GoroutineRecord) GetAddress() uint64 {
	return r.Address
}

func (r *GoroutineRecord) Read(rd *bufio.Reader) error {
	var err error
	r.Address, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.StackPtr, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.RoutineID, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.CreatorPtr, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	status, err := binary.ReadUvarint(rd)
	if err != nil {
		return err
	}
	r.Status = Status(status)

	systemInt, err := binary.ReadUvarint(rd)
	if err != nil {
		return err
	}
	r.System = (systemInt != 0)

	bgInt, err := binary.ReadUvarint(rd)
	if err != nil {
		return err
	}
	r.Background = (bgInt != 0)

	r.WaitTime, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	contentSize, err := binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	waitReason := make([]byte, contentSize)
	_, err = io.ReadFull(rd, waitReason)
	if err != nil {
		return err
	}

	r.WaitReason = string(waitReason)

	r.ContextPtr, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.ThreadAddress, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.TopDefer, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.TopPanic, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	return nil
}

func (r *GoroutineRecord) Repr() string {
	if r.Status == Waiting {
		format := "Goroutine at address 0x%x (routine ID = %d, status = %s, wait reason = %s)"
		return fmt.Sprintf(format, r.Address, r.RoutineID, r.Status.Repr(), r.WaitReason)
	}

	format := "Goroutine at address 0x%x (routine ID = %d, status = %s, stack address = 0x%x)"
	return fmt.Sprintf(format, r.Address, r.RoutineID, r.Status.Repr(), r.StackPtr)
}

// Possible statuses:
// 0 = idle
// 1 = runnable
// 3 = syscall
// 4 = waiting
type Status uint64

const (
	Idle Status = iota
	Runnable
	Syscall
	Waiting
)

func (s Status) Repr() string {
	switch s {
	case Idle:
		return "Idle"
	case Runnable:
		return "Runnable"
	case Syscall:
		return "Syscall"
	case Waiting:
		return "Waiting"
	default:
		return "Unknown"
	}
}
