package mbgl

import (
	"testing"
)

func TestNewFileSource(t *testing.T) {

    fileSource := NewDefaultFileSource("testdata/cache.sqlite", ".")
	defer fileSource.Destroy()

	if &fileSource.cptr == nil {
		t.Fatal("NewDefaultFileSource returned nil")
	}	
}
