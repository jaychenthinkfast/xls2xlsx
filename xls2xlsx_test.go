package xls2xlsx

import (
	"os"
	"testing"
)

func TestConvert(t *testing.T) {
	Convert("example/test.xls")
	if _, err := os.Stat("example/test.xlsx"); err != nil {
		t.Error(err)
	}
}
