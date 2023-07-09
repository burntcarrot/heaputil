package record

const Header = "go1.7 heap dump\n"

type RecordType int

// Set up record type constants.
const (
	EOF RecordType = iota
	Object
	OtherRoot
	TypeDescriptor
	Goroutine
	StackFrame
	DumpParams
	RegisteredFinalizer
	ITab
	OSThread
	MemStats
	QueuedFinalizer
	DataSegment
	BSSSegment
	DeferRecord
	PanicRecord
	AllocFreeProfile
	AllocStackTraceSample
)
