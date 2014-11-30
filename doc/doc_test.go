package doc

import (
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	docs := LoadConfig()
	if docs == nil {
		t.Error("errors")
	}
	fmt.Printf("%#v", docs)
}
