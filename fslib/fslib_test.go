package fslib

import (
	"testing"
)

type testDirEntry struct {
	name  string
	isdir bool
}

func (t testDirEntry) Name() string {
	return t.name
}

func (t testDirEntry) IsDir() bool {
	return t.isdir
}

func TestCreateStrComp(t *testing.T) {
	fn := CreateStringComp("test")
	dire := testDirEntry{"test", false}
	got := fn(dire)
	if !got {
		t.Errorf("comparison failed")
	}
}
