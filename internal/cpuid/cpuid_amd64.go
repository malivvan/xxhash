//go:build amd64 && !appengine && !noasm && gc && !purego

package cpuid

func getCpuid(op, sub uint32) (eax, ebx, ecx, edx uint32)
func xgetbv0() (eax, edx uint32)

func detect() Info {
	var i Info
	maxLeaf, _, _, _ := getCpuid(0, 0)
	if maxLeaf < 1 {
		return i
	}
	_, _, ecx1, edx1 := getCpuid(1, 0)
	i.set(SSE2, edx1&(1<<26) != 0)
	i.set(POPCNT, ecx1&(1<<23) != 0)

	var xcr0 uint32
	xsave := false
	if ecx1&(1<<26) != 0 && ecx1&(1<<27) != 0 {
		xcr0, _ = xgetbv0()
		xsave = true
	}
	avxState := xsave && xcr0&0x6 == 0x6
	avx512State := avxState && (xcr0>>5)&7 == 7

	if maxLeaf < 7 {
		return i
	}
	_, ebx7, ecx7, _ := getCpuid(7, 0)
	i.set(BMI2, ebx7&(1<<8) != 0)
	i.set(AVX2, avxState && ebx7&(1<<5) != 0)

	if !avx512State || ebx7&(1<<16) == 0 {
		return i
	}
	i.set(AVX512F, true)
	i.set(AVX512BW, ebx7&(1<<30) != 0)
	i.set(AVX512VL, ebx7&(1<<31) != 0)
	i.set(AVX512VBMI, ecx7&(1<<1) != 0)
	i.set(AVX512VBMI2, ecx7&(1<<6) != 0)
	i.set(AVX512BITALG, ecx7&(1<<12) != 0)
	i.set(AVX512VPOPCNTDQ, ecx7&(1<<14) != 0)
	i.set(GFNI, ecx7&(1<<8) != 0)
	return i
}
