package resenum

import (
	"fmt"
	"strconv"
	"strings"
)

type Pair struct {
	X uint
	Y uint
}

// parses `s`, assuming it is formatteda as `x:y`. Returns 0,0,error on failure
func ParseUintPair(s string) (Pair, error) {
	var rx uint64
	var ry uint64
	var err error
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		goto bail
	}

	rx, err = strconv.ParseUint(parts[0], 10, 32)
	if err != nil {
		goto bail
	}

	ry, err = strconv.ParseUint(parts[1], 10, 32)
	if err != nil {
		goto bail
	}

	return Pair{uint(rx), uint(ry)}, nil
bail:
	return Pair{0, 0}, fmt.Errorf("pair could not be parsed: `%s`", s)
}

// parses `s`, assuming it is formatteda as `x:y`. Returns 0,0,error on failure
func ParseOptionalUintPair(s string, def Pair) (Pair, error) {
	var rx uint64
	var ry uint64
	var err error
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		goto bail
	}

	rx, err = strconv.ParseUint(parts[0], 10, 32)
	if err != nil {
		rx = uint64(def.X)
	}

	ry, err = strconv.ParseUint(parts[1], 10, 32)
	if err != nil {
		ry = uint64(def.Y)
	}

	return Pair{uint(rx), uint(ry)}, nil
bail:
	return Pair{0, 0}, fmt.Errorf("pair could not be parsed: `%s`", s)
}
