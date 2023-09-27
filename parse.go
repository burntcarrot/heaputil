package heaputil

import (
	"bufio"
	"fmt"

	"github.com/burntcarrot/heaputil/record"
)

type RecordData struct {
	RecordType  record.RecordType
	Repr        string
	HasPointers bool
	RowID       string
	Pointers    []PointerData
}

type PointerData struct {
	Index    int
	Address  string
	Incoming string
	Outgoing string
}

func ParseDump(rd *bufio.Reader) ([]RecordData, error) {
	err := record.ReadHeader(rd)
	if err != nil {
		return nil, err
	}

	var dumpParams *record.DumpParamsRecord
	records := []RecordData{}

	for {
		r, err := record.ReadRecord(rd)
		if err != nil {
			return nil, err
		}

		_, isEOF := r.(*record.EOFRecord)
		if isEOF {
			break
		}

		dp, isDumpParams := r.(*record.DumpParamsRecord)
		if isDumpParams {
			dumpParams = dp
		}

		recordInfo := RecordData{
			RecordType: record.GetRecordType(r),
			Repr:       r.Repr(),
		}

		p, isParent := r.(record.ParentGuard)
		if isParent {
			incoming, outgoing := record.ParsePointers(p, dumpParams)
			for i := 0; i < len(outgoing); i++ {
				if outgoing[i] != 0 {
					a, hasAddress := r.(record.AddressGuard)
					if hasAddress {
						address := a.GetAddress() + p.GetFields()[i]
						pointerInfo := PointerData{
							Index:    i,
							Address:  fmt.Sprintf("%x", address),
							Incoming: fmt.Sprintf("%x", incoming[i]),
							Outgoing: fmt.Sprintf("%x", outgoing[i]),
						}
						recordInfo.Pointers = append(recordInfo.Pointers, pointerInfo)
						recordInfo.HasPointers = true
					}
				}
			}
		}

		recordInfo.RowID = fmt.Sprintf("row%d", len(records)+1)

		records = append(records, recordInfo)
	}

	return records, nil
}

// PrintDump prints the heap dump information to stdout.
// CAUTION: can be too verbose!
func PrintDump(rd *bufio.Reader) error {
	err := record.ReadHeader(rd)
	if err != nil {
		return err
	}

	var dumpParams *record.DumpParamsRecord

	for {
		r, err := record.ReadRecord(rd)
		if err != nil {
			return err
		}

		_, isEOF := r.(*record.EOFRecord)
		if isEOF {
			break
		}

		dp, isDumpParams := r.(*record.DumpParamsRecord)
		if isDumpParams {
			dumpParams = dp
		}

		// Print pointer information.
		p, isParent := r.(record.ParentGuard)
		if isParent {
			incoming, outgoing := record.ParsePointers(p, dumpParams)
			for i := 0; i < len(outgoing); i++ {
				// If outgoing (destination) is valid, then print it.
				if outgoing[i] != 0 {
					a, hasAddress := r.(record.AddressGuard)
					if hasAddress {
						address := a.GetAddress() + p.GetFields()[i]
						format := "\tPointer(#%d) at address 0x%x (incoming = 0x%x, outgoing = 0x%x)\n"
						fmt.Printf(format, i, address, incoming[i], outgoing[i])
					}
				}
			}
		}

		// Display record.
		fmt.Println(r.Repr())
	}

	return nil
}
