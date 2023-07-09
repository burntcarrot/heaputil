package record

import (
	"bufio"
	"encoding/binary"
	"fmt"
)

// Each record should accept a reader and read from it, and should also have a representation (Repr) of itself.
type Record interface {
	Read(r *bufio.Reader) error
	Repr() string
}

// ReadRecord reads a record by checking the record type (represented by a uvarint-encoded integer).
func ReadRecord(r *bufio.Reader) (Record, error) {
	// Read record type
	recordType, err := binary.ReadUvarint(r)
	if err != nil {
		return nil, err
	}

	var record Record

	// Follow the order as described in the heap dump format.
	// 	Each record starts with a uvarint-encoded integer describing the type of the record:

	// 0 = EOF
	// 1 = object
	// 2 = otherroot
	// 3 = type
	// 4 = goroutine
	// 5 = stack frame
	// 6 = dump params
	// 7 = registered finalizer
	// 8 = itab
	// 9 = OS thread
	// 10 = mem stats
	// 11 = queued finalizer
	// 12 = data segment
	// 13 = bss segment
	// 14 = defer record
	// 15 = panic record
	// 16 = alloc/free profile record
	// 17 = alloc stack trace sample
	switch RecordType(recordType) {
	case EOF:
		record = &EOFRecord{}
	case Object:
		record = &ObjectRecord{}
	case OtherRoot:
		record = &OtherRootRecord{}
	case TypeDescriptor:
		record = &TypeDescriptorRecord{}
	case Goroutine:
		record = &GoroutineRecord{}
	case StackFrame:
		record = &StackFrameRecord{}
	case DumpParams:
		record = &DumpParamsRecord{}
	case RegisteredFinalizer:
		record = &RegisteredFinalizerRecord{}
	case ITab:
		record = &ITabRecord{}
	case OSThread:
		record = &OSThreadRecord{}
	case MemStats:
		record = &MemStatsRecord{}
	case QueuedFinalizer:
		record = &QueuedFinalizerRecord{}
	case DataSegment:
		record = &DataSegmentRecord{}
	case BSSSegment:
		record = &BSSSegmentRecord{}
	case DeferRecord:
		record = &DeferRecordRecord{}
	case PanicRecord:
		record = &PanicRecordRecord{}
	case AllocFreeProfile:
		record = &AllocFreeProfileRecord{}
	case AllocStackTraceSample:
		record = &AllocStackTraceSampleRecord{}
	default:
		return nil, fmt.Errorf("unexpected record type: %v", recordType)
	}

	// Read record. Each record implements the Record interface.
	err = record.Read(r)
	if err != nil {
		return nil, err
	}

	return record, nil
}

// AddressGuard is used as a "guard" for checking if a record has address.
type AddressGuard interface {
	GetAddress() uint64
}

// ParentGuard is used as a "guard" for checking the parent pointer information.
type ParentGuard interface {
	GetAddress() uint64
	GetContent() []byte
	GetFields() []uint64
}
