package record

import (
	"bufio"
	"fmt"
	"io"
)

func ReadHeader(r *bufio.Reader) error {
	h := make([]byte, len(Header))
	n, err := io.ReadFull(r, h)
	if err != nil {
		return err
	}

	if n != len(Header) {
		return fmt.Errorf("header length doesn't match, expected=%d, got=%d", len(Header), n)
	}

	if string(h) != Header {
		return fmt.Errorf("header value doesn't match, got = %s", string(h))
	}
	return nil
}
