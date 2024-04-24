package memory

// This type represents the GBA General Internal Memory described on [GBATEK].
//
// The unused memory won't be mapped and so not represented in the underlying
// type.
//
// [GBATEK]: https://problemkaputt.de/gbatek.htm#gbamemorymap
type Internal struct {
	// BIOS - System ROM        -  16 kbytes
	Bios [16 * Kilo]byte
	// WRAM - On-board Work RAM - 256 kbytes
	BoardWRAM [256 * Kilo]byte
	// WRAM - Ob-chip Work RAM  -  32 kbytes
	ChipWRAM [32 * Kilo]byte
	// IO Registers             -   2 kbytes
	IORegisters [2 * Kilo]byte
}
