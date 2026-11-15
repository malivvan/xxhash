package cpuid

import "strconv"

type Feature uint8

const (
	SSE2 Feature = iota
	POPCNT
	BMI2
	AVX2
	AVX512F
	AVX512BW
	AVX512VL
	AVX512VBMI
	AVX512VBMI2
	AVX512BITALG
	AVX512VPOPCNTDQ
	GFNI

	numFeatures
)

var featureNames = [numFeatures]string{
	SSE2:            "SSE2",
	POPCNT:          "POPCNT",
	BMI2:            "BMI2",
	AVX2:            "AVX2",
	AVX512F:         "AVX512F",
	AVX512BW:        "AVX512BW",
	AVX512VL:        "AVX512VL",
	AVX512VBMI:      "AVX512VBMI",
	AVX512VBMI2:     "AVX512VBMI2",
	AVX512BITALG:    "AVX512BITALG",
	AVX512VPOPCNTDQ: "AVX512VPOPCNTDQ",
	GFNI:            "GFNI",
}

func (f Feature) String() string {
	if f >= numFeatures {
		return "Feature(" + strconv.Itoa(int(f)) + ")"
	}
	return featureNames[f]
}

type Info struct {
	bits uint32
}

func (i Info) Has(f Feature) bool {
	return i.bits&(1<<f) != 0
}

func (i Info) Supports(f ...Feature) bool {
	for _, want := range f {
		if !i.Has(want) {
			return false
		}
	}
	return true
}

func (i *Info) Disable(f ...Feature) {
	for _, drop := range f {
		i.bits &^= 1 << drop
	}
}

func (i *Info) set(f Feature, ok bool) {
	if ok {
		i.bits |= 1 << f
	}
}

var CPU = detect()
