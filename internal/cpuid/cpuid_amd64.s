

//go:build amd64 && !appengine && !noasm && gc && !purego

#include "textflag.h"

// func getCpuid(op, sub uint32) (eax, ebx, ecx, edx uint32)
TEXT ·getCpuid(SB), NOSPLIT, $0-24
	MOVL op+0(FP), AX
	MOVL sub+4(FP), CX
	CPUID
	MOVL AX, eax+8(FP)
	MOVL BX, ebx+12(FP)
	MOVL CX, ecx+16(FP)
	MOVL DX, edx+20(FP)
	RET

// func xgetbv0() (eax, edx uint32)
//
// XGETBV reads XCR0 selected by ECX. The instruction is encoded as raw bytes
// because the assembler has no mnemonic for it.
TEXT ·xgetbv0(SB), NOSPLIT, $0-8
	XORL CX, CX
	BYTE $0x0f; BYTE $0x01; BYTE $0xd0 // XGETBV
	MOVL AX, eax+0(FP)
	MOVL DX, edx+4(FP)
	RET
