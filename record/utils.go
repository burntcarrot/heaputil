package record

import (
	"bufio"
	"encoding/binary"
	"fmt"
)

// ReadFieldList is a utility function to read the fieldlist primitive present in the heap dump.
func ReadFieldList(rd *bufio.Reader) ([]uint64, error) {
	// A field list consists of repeated pairs of uvarints encoding a field kind and a field offset, followed by an end-of-list marker. The only possible kind is 1=Ptr.
	// Earlier versions of the heap dump could contain 2=Iface and 3=Eface, but the runtime no longer tracks that information, so it is not present in the dump. Interface values appear as a pair of pointers. 0=Eol is the end of list marker. The end of list marker does not have a corresponding offset.
	fields := []uint64{}

	for {
		kind, err := binary.ReadUvarint(rd)
		if err != nil {
			return nil, err
		}
		if kind == 0 { // EOL marker
			break
		}

		value, err := binary.ReadUvarint(rd)
		if err != nil {
			return nil, err
		}
		if kind == 0 { // EOL marker
			break
		}

		// Append field offset.
		fields = append(fields, value)
	}

	return fields, nil
}

// ParsePointers is a utility function for parsing the pointers of a parent "node"/object.
func ParsePointers(p ParentGuard, dp *DumpParamsRecord) ([]uint64, []uint64) {
	var endian binary.ByteOrder
	endian = binary.LittleEndian
	if dp.BigEndian {
		endian = binary.BigEndian
	}

	fields := p.GetFields()
	contents := p.GetContent()

	incoming := make([]uint64, len(fields))
	outgoing := make([]uint64, len(fields))

	for i := 0; i < len(fields); i++ {
		fieldOffset := fields[i]
		incoming[i] = p.GetAddress() + fieldOffset
		switch dp.PtrSize {
		case 2:
			outgoing[i] = uint64(endian.Uint16(contents[fieldOffset:]))
		case 4:
			outgoing[i] = uint64(endian.Uint32(contents[fieldOffset:]))
		case 8:
			outgoing[i] = endian.Uint64(contents[fieldOffset:])
		default:
			panic(fmt.Sprintf("unexpected pointer size %d", dp.PtrSize))
		}
	}

	return incoming, outgoing
}

func GetRecordType(record Record) RecordType {
	switch r := record.(type) {
	case *EOFRecord:
		return EOF
	case *ObjectRecord:
		return Object
	case *OtherRootRecord:
		return OtherRoot
	case *TypeDescriptorRecord:
		return TypeDescriptor
	case *GoroutineRecord:
		return Goroutine
	case *StackFrameRecord:
		return StackFrame
	case *DumpParamsRecord:
		return DumpParams
	case *RegisteredFinalizerRecord:
		return RegisteredFinalizer
	case *ITabRecord:
		return ITab
	case *OSThreadRecord:
		return OSThread
	case *MemStatsRecord:
		return MemStats
	case *QueuedFinalizerRecord:
		return QueuedFinalizer
	case *DataSegmentRecord:
		return DataSegment
	case *BSSSegmentRecord:
		return BSSSegment
	case *DeferRecordRecord:
		return DeferRecord
	case *PanicRecordRecord:
		return PanicRecord
	case *AllocFreeProfileRecord:
		return AllocFreeProfile
	case *AllocStackTraceSampleRecord:
		return AllocStackTraceSample
	default:
		fmt.Printf("cannot find type: %T\n", r)
		return -1
	}
}

func GetRecordTypeStr(record Record) string {
	switch r := record.(type) {
	case *EOFRecord:
		return "EOF"
	case *ObjectRecord:
		return "Object"
	case *OtherRootRecord:
		return "OtherRoot"
	case *TypeDescriptorRecord:
		return "TypeDescriptor"
	case *GoroutineRecord:
		return "Goroutine"
	case *StackFrameRecord:
		return "StackFrame"
	case *DumpParamsRecord:
		return "DumpParams"
	case *RegisteredFinalizerRecord:
		return "RegisteredFinalizer"
	case *ITabRecord:
		return "ITab"
	case *OSThreadRecord:
		return "OSThread"
	case *MemStatsRecord:
		return "MemStats"
	case *QueuedFinalizerRecord:
		return "QueuedFinalizer"
	case *DataSegmentRecord:
		return "DataSegment"
	case *BSSSegmentRecord:
		return "BSSSegment"
	case *DeferRecordRecord:
		return "DeferRecord"
	case *PanicRecordRecord:
		return "PanicRecord"
	case *AllocFreeProfileRecord:
		return "AllocFreeProfile"
	case *AllocStackTraceSampleRecord:
		return "AllocStackTraceSample"
	default:
		fmt.Printf("cannot find type: %T\n", r)
		return "Unknown"
	}
}

func GetStrFromRecordType(rType RecordType) string {
	switch rType {
	case EOF:
		return "EOF"
	case Object:
		return "Object"
	case OtherRoot:
		return "OtherRoot"
	case TypeDescriptor:
		return "TypeDescriptor"
	case Goroutine:
		return "Goroutine"
	case StackFrame:
		return "StackFrame"
	case DumpParams:
		return "DumpParams"
	case RegisteredFinalizer:
		return "RegisteredFinalizer"
	case ITab:
		return "ITab"
	case OSThread:
		return "OSThread"
	case MemStats:
		return "MemStats"
	case QueuedFinalizer:
		return "QueuedFinalizer"
	case DataSegment:
		return "DataSegment"
	case BSSSegment:
		return "BSSSegment"
	case DeferRecord:
		return "DeferRecord"
	case PanicRecord:
		return "PanicRecord"
	case AllocFreeProfile:
		return "AllocFreeProfile"
	case AllocStackTraceSample:
		return "AllocStackTraceSample"
	default:
		return "Unknown"
	}
}
