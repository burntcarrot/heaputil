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
