"".(*multiplyRealNode).evalReal STEXT size=193 args=0x28 locals=0x40
	0x0000 00000 (main.go:27)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:27)	CMPQ	SP, 16(CX)
	0x000d 00013 (main.go:27)	JLS	183
	0x0013 00019 (main.go:27)	SUBQ	$64, SP
	0x0017 00023 (main.go:27)	MOVQ	BP, 56(SP)
	0x001c 00028 (main.go:27)	LEAQ	56(SP), BP
	0x0021 00033 (main.go:28)	MOVQ	"".node+72(SP), AX
	0x0026 00038 (main.go:28)	MOVQ	(AX), CX
	0x0029 00041 (main.go:28)	MOVQ	8(AX), DX
	0x002d 00045 (main.go:28)	MOVQ	24(CX), CX
	0x0031 00049 (main.go:28)	MOVQ	DX, (SP)
	0x0035 00053 (main.go:28)	MOVQ	"".row+80(SP), DX
	0x003a 00058 (main.go:28)	MOVQ	DX, 8(SP)
	0x003f 00063 (main.go:28)	MOVQ	"".row+88(SP), BX
	0x0044 00068 (main.go:28)	MOVQ	BX, 16(SP)
	0x0049 00073 (main.go:28)	CALL	CX
	0x004b 00075 (main.go:29)	MOVQ	"".node+72(SP), AX
	0x0050 00080 (main.go:29)	MOVQ	16(AX), CX
	0x0054 00084 (main.go:28)	MOVSD	24(SP), X0
	0x005a 00090 (main.go:28)	MOVSD	X0, "".v1+48(SP)
	0x0060 00096 (main.go:28)	MOVBLZX	32(SP), DX
	0x0065 00101 (main.go:28)	MOVB	DL, "".null1+47(SP)
	0x0069 00105 (main.go:29)	MOVQ	24(AX), AX
	0x006d 00109 (main.go:29)	MOVQ	24(CX), CX
	0x0071 00113 (main.go:29)	MOVQ	AX, (SP)
	0x0075 00117 (main.go:29)	MOVQ	"".row+80(SP), AX
	0x007a 00122 (main.go:29)	MOVQ	AX, 8(SP)
	0x007f 00127 (main.go:29)	MOVQ	"".row+88(SP), AX
	0x0084 00132 (main.go:29)	MOVQ	AX, 16(SP)
	0x0089 00137 (main.go:29)	CALL	CX
	0x008b 00139 (main.go:29)	MOVBLZX	32(SP), AX
	0x0090 00144 (main.go:30)	MOVSD	"".v1+48(SP), X0
	0x0096 00150 (main.go:30)	MULSD	24(SP), X0
	0x009c 00156 (main.go:30)	MOVSD	X0, "".~r1+96(SP)
	0x00a2 00162 (main.go:30)	MOVBLZX	"".null1+47(SP), CX
	0x00a7 00167 (main.go:30)	ORL	AX, CX
	0x00a9 00169 (main.go:30)	MOVB	CL, "".~r2+104(SP)
	0x00ad 00173 (main.go:30)	MOVQ	56(SP), BP
	0x00b2 00178 (main.go:30)	ADDQ	$64, SP
	0x00b6 00182 (main.go:30)	RET
	0x00b7 00183 (main.go:30)	NOP
	0x00b7 00183 (main.go:27)	CALL	runtime.morestack_noctxt(SB)
	0x00bc 00188 (main.go:27)	JMP	0