//go:build amd64 && !appengine && !noasm && gc && !purego

package cpuid

import "testing"

func TestDetectAMD64(t *testing.T) {
	if !CPU.Has(SSE2) {
		t.Fatal("SSE2 not detected on amd64")
	}
	if CPU.Has(AVX512F) {
		if !CPU.Has(AVX2) {
			t.Fatal("AVX-512F without AVX2")
		}
		if !CPU.Has(AVX512BW) || !CPU.Has(AVX512VL) {
			t.Log("AVX-512 without BW/VL: kernels stay on the AVX2 tier")
		}
	}
	if CPU.Has(GFNI) && !CPU.Has(AVX512F) {
		t.Fatal("GFNI reported without the AVX-512 form")
	}
}
