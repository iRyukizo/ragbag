package memory

// This type represents the GBA External Memory (Game Pak) described on
// [GBATEK].
//
// The unused memory won't be mapped and so not represented in the underlying
// type.
//
// [GBATEK]: https://problemkaputt.de/gbatek.htm#gbamemorymap
type External struct {
	// Game Pak ROM/FlashROM - Wait State 0   - Maximum 32 mbytes
	GPROM0 [32 * Mega]byte
	// Game Pak ROM/FlashROM - Wait State 1   - Maximum 32 mbytes
	GPROM1 [32 * Mega]byte
	// Game Pak ROM/FlashROM - Wait State 2   - Maximum 32 mbytes
	GPROM2 [32 * Mega]byte
	// Game Pak SRAM         - 8bit Bus width - Maximum 64 kbytes
	GPSRAM [64 * Kilo]byte
}
