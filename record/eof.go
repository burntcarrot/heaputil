package record

import "bufio"

type EOFRecord struct{}

func (r *EOFRecord) Read(rd *bufio.Reader) error {
	return nil
}

func (r *EOFRecord) Repr() string {
	return "EOF"
}
