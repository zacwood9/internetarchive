package internetarchive

import (
	"testing"
)

func TestItem_GetMetadata(t *testing.T) {
	item := Item{"gd1977-05-08.mtx.dan.29511.flac16",
		471380,
		"1977-05-08",
		4.59,
		"",
		"",
		metadata{}}

	if err := item.GetMetadata(); err != nil {
		t.Fatal(err)
	}
}
