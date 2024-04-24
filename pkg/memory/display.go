package memory

// This type represents the GBA Internal Display Memory described on [GBATEK].
//
// The unused memory won't be mapped and so not represented in the underlying
// type.
//
// [GBATEK]: https://problemkaputt.de/gbatek.htm#gbamemorymap
type Display struct {
	// BG/OBJ Palette RAM   - 1  kbyte
	PaletteRAM [1 * Kilo]byte
	// VRAM - Video RAM     - 96 kbytes
	VideoRAM [96 * Kilo]byte
	// OAM - OBJ Attributes - 1  kbyte
	OAM [1 * Kilo]byte
}
