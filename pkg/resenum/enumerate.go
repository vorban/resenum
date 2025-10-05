package resenum

import "fmt"

type ResolutionType int

func (rt ResolutionType) ToString() string {
	switch rt {
	case ResSD:
		return "SD"
	case ResHD:
		return "HD"
	case ResFHD:
		return "Full HD"
	case ResQHD:
		return "QHD"
	case ResUHD:
		return "4K UHD"
	default:
		return ""
	}
}

const (
	ResNonStandard ResolutionType = iota
	ResSD
	ResHD
	ResFHD
	ResQHD
	ResUHD
)

type Resolution struct {
	x    uint
	y    uint
	Type ResolutionType
}

func (res Resolution) ToString() string {
	return fmt.Sprintf("%dx%d", res.x, res.y)
}

func (res Resolution) AsPair() Pair {
	return Pair{res.x, res.y}
}

func Enumerate(ratio Pair, min Pair, max Pair) []Resolution {
	var resolutions []Resolution

	x, y := min.X, min.Y
	for x*ratio.Y != y*ratio.X {
		if x*ratio.Y < y*ratio.X {
			x++
		} else {
			y++
		}
	}

	k := getMultiplier(x, ratio.X)
	for ; x <= max.X && y <= max.Y; x, y = x+ratio.X, y+ratio.Y {
		res_type := getResolutionType(k)
		k++

		if x == 0 || y == 0 {
			continue
		}

		resolutions = append(resolutions, Resolution{x, y, res_type})
	}

	return resolutions
}

func getMultiplier(x uint, r uint) uint {
	if r == 0 {
		return 0
	}

	return x / r
}

func getResolutionType(k uint) ResolutionType {
	switch k {
	case 40:
		return ResSD
	case 80:
		return ResHD
	case 120:
		return ResFHD
	case 160:
		return ResQHD
	case 240:
		return ResUHD
	default:
		return ResNonStandard
	}
}
