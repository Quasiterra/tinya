package core

import (
	"testing"
)

func TestListBuckets(t *testing.T) {
	if _, err := ListBuckets(); err != nil {
		t.Error(err)
	}
}
