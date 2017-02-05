package read

import (
	"io/ioutil"
	"strconv"
	"strings"
)

const tempFilePath = "/sys/class/thermal/thermal_zone0/temp"

// Temp reads cpu temperature from tempFilePath
// and returns in degrees
func Temp() (res float64, err error) {
	var rawTemp []byte
	rawTemp, err = ioutil.ReadFile(tempFilePath)
	if err != nil {
		return
	}
	res, err = strconv.ParseFloat(strings.TrimSpace(string(rawTemp)), 64)
	res /= float64(1000)
	return
}