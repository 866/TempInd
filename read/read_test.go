package read

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

const (
	TEMPHIGHRANGE = float64(79)
	TEMPLOWRANGE  = float64(15)
	STRESSDURATION time.Duration = time.Second * 2
)

// TestTemp provides testing read.Temp function
func TestTemp(t *testing.T) {
	log.Println("Check whether Temp() function outputs the proper temperature.")
	val, err := Temp()
	if err != nil {
		t.Error("Error occured: ", err)
	}
	if val < TEMPLOWRANGE || val > TEMPHIGHRANGE {
		t.Error("Expected temperature to be in range: ", TEMPLOWRANGE,
			"-", TEMPHIGHRANGE, "\nGot: ", val)
	}
}

func stress(duration time.Duration) {
	a := rand.Float64()
	b := rand.Float64() 
Cycle:
	for {
		select {
		case <-time.After(STRESSDURATION):
			break Cycle
		default:
			// Some dummy operation to load the CPU
			a = a * b	
		}
	}
}

// TestStress provides the stress on CPU and measures whether the temperature increased
// May not work on some powerful machines however does work on my Raspbery PI3 :P
func TestSress(t *testing.T) {
	log.Println("Check whether Temp() function can track temperature change caused by stress.")
	var (
		valFirst, valSecond float64
		err error
	)
	valFirst, err = Temp()
	if err != nil {
		t.Error("Error occured: ", err)
	}
	for i := 0; i < 5; i++ {
		go stress(STRESSDURATION)
	}
	time.Sleep(STRESSDURATION)
	valSecond, err = Temp()
	if err != nil {
		t.Error("Error occured: ", err)
	}
	if valFirst > valSecond {
		t.Error("Expected the temperature to be higher after the stress. Got:\nbefore the stress: ",
			valFirst, "\nafter the stress: ", valSecond)
	}
}

