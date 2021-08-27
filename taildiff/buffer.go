package main

import (
	"bytes"
)

// ChangeDetectingBuffer is a reusable io.Writer that detect changes across
// consecutive batchs of writes. A batch starts by Rewind and ends by Done.
// A new ChangeDetectingBuffer does not need to call Rewind before its first
// use. Only check the Changedstatus after ending the batch by calling Done.
type ChangeDetectingBuffer struct {
	bytes   []byte
	off     int
	Changed bool
}

// Done ends the batch of writes, updating the Changed status and truncating
// the underlying byte slice if necessary.
func (b *ChangeDetectingBuffer) Done() {
	if b.off != len(b.bytes) {
		b.Changed = true
	}
	b.bytes = b.bytes[:b.off]
}

// Rewind resets the Changed status and prepares for a new batch of writes.
func (b *ChangeDetectingBuffer) Rewind() {
	b.off = 0
	b.Changed = false
}

// Write lets ChangeDetectingBuffer satisfy io.Writer.
func (b *ChangeDetectingBuffer) Write(p []byte) (int, error) {
	if len(b.bytes[b.off:]) < len(p) {
		b.Changed = true
	} else if !bytes.Equal(b.bytes[b.off:b.off+len(p)], p) {
		b.Changed = true
	}

	l := copy(b.bytes[b.off:], p)
	b.bytes = append(b.bytes, p[l:]...)
	b.off += len(p)
	return len(p), nil
}

// String lets ChangeDetectingBuffer satisfy fmt.Stringer
func (b *ChangeDetectingBuffer) String() string {
	return string(b.bytes)
}
