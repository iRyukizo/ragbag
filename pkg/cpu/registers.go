package cpu

// This type represent the GBA CPU register set describer on [GBATEK].
//
// There's a total of 37 32-bit registers, 31 general registers and 6 status
// registers.
//
// [GBATEK]: https://problemkaputt.de/gbatek.htm#armcpuregisterset
type RegisterSet struct {
	// General registers
	// R0-R12 - General Purpose Registers
	GPRegisters [13]uint32
	// R8-R12 - FIQ Registers
	FIQRegisters [5]uint32
	// R13 - Stack Pointer
	Registers13 [6]uint32
	// R14 - Link Register
	Registers14 [6]uint32
	// R15 - Program Counter
	Register15 uint32
	// Status registers
	// CPSR
	CPSR uint32
	// SPSR
	SPSR [5]uint32
}
