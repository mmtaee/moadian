package tax_id

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

func GenerateTaxID(fiscal string, serial int, timestamp int64) string {
	tPad, tHex := timeNormalizer(timestamp)
	sPad, sHex := serialNormalizer(serial)
	verhoeffChSum := verhoeffCheckSum(fiscal, sPad, tPad)
	return fiscal + tHex + sHex + verhoeffChSum
}

// function create remain days from 1970 padded and hex padded
func timeNormalizer(timestamp int64) (string, string) {
	var t int
	if timestamp == 0 {
		t = remainingDays(time.Now().Unix())
	} else {
		valid := timestampValidator(timestamp)
		if !valid {
			log.Fatalln("invalid timestamp")
		}
		t = remainingDays(timestamp)
	}
	return fmt.Sprintf("%06d", t), fmt.Sprintf("%05s", strings.ToUpper(strconv.FormatInt(int64(t), 16)))
}

func timestampValidator(timestamp int64) bool {
	if timestamp < 0 {
		return false
	}
	t := time.Unix(timestamp, 0)
	return t.Unix() == timestamp
}

func remainingDays(i int64) int {
	return int(math.Floor(float64(i / 86400)))
}

// function return serial number with len 12 in string and hex
func serialNormalizer(serial int) (string, string) {
	sPad := fmt.Sprintf("%012d", serial)
	i, _ := strconv.Atoi(sPad)
	sHex := fmt.Sprintf("%010X", i)
	return sPad, sHex
}
