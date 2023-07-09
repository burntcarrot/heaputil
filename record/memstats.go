package record

import (
	"bufio"
	"encoding/binary"
	"fmt"
)

// uvarint: Alloc
// uvarint: TotalAlloc
// uvarint: Sys
// uvarint: Lookups
// uvarint: Mallocs
// uvarint: Frees
// uvarint: HeapAlloc
// uvarint: HeapSys
// uvarint: HeapIdle
// uvarint: HeapInuse
// uvarint: HeapReleased
// uvarint: HeapObjects
// uvarint: StackInuse
// uvarint: StackSys
// uvarint: MSpanInuse
// uvarint: MSpanSys
// uvarint: MCacheInuse
// uvarint: MCacheSys
// uvarint: BuckHashSys
// uvarint: GCSys
// uvarint: OtherSys
// uvarint: NextGC
// uvarint: LastGC
// uvarint: PauseTotalNs
// 256 uvarints: PauseNs
// uvarint: NumGC
type MemStatsRecord struct {
	Alloc        uint64
	TotalAlloc   uint64
	Sys          uint64
	Lookups      uint64
	Mallocs      uint64
	Frees        uint64
	HeapAlloc    uint64
	HeapSys      uint64
	HeapIdle     uint64
	HeapInUse    uint64
	HeapReleased uint64
	HeapObjects  uint64
	StackInUse   uint64
	StackSys     uint64
	MSpanInUse   uint64
	MSpanSys     uint64
	MCacheInUse  uint64
	MCacheSys    uint64
	BuckHashSys  uint64
	GCSys        uint64
	OtherSys     uint64
	NextGC       uint64
	LastGC       uint64
	PauseTotalNs uint64
	PauseNs      [256]uint64
	NumGC        uint64
}

func (r *MemStatsRecord) Read(rd *bufio.Reader) error {
	var err error

	r.Alloc, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.TotalAlloc, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.Sys, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.Lookups, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.Mallocs, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.Frees, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.HeapAlloc, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.HeapSys, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.HeapIdle, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.HeapInUse, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.HeapReleased, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.HeapObjects, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.StackInUse, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.StackSys, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.MSpanInUse, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.MSpanSys, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.MCacheInUse, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.MCacheSys, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.BuckHashSys, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.GCSys, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.OtherSys, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.NextGC, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.LastGC, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	r.PauseTotalNs, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	for i := 0; i < 256; i++ {
		r.PauseNs[i], err = binary.ReadUvarint(rd)
		if err != nil {
			return err
		}
	}

	r.NumGC, err = binary.ReadUvarint(rd)
	if err != nil {
		return err
	}

	return nil
}

func (r *MemStatsRecord) Repr() string {
	format := "MemStats: %+v"

	return fmt.Sprintf(format, *r)
}
