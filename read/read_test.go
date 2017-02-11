package read

import (
	"testing"
)

const (
	TEMPHIGHRANGE = float64(79)
	TEMPLOWRANGE  = float64(15)
)

// TestTemp provides testing read.Temp function
func TestTemp(t *testing.T) {
	val, err := Temp()
	if err != nil {
		t.Error("Error occured: ", err)
	}
	if val < TEMPLOWRANGE || val > TEMPHIGHRANGE {
		t.Error("Expected temperature to be in range: ", TEMPLOWRANGE,
			"-", TEMPHIGHRANGE, "\nGot: ", val)
	}
}
