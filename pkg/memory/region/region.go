package region

// Those const represents the GBA Memory Map.
const (
	// General Internal Memory
	BIOSBegin        uint = 0x00000000
	BIOSEnd               = 0x00003FFF
	Unused1Begin          = 0x00004000
	Unused1End            = 0x01FFFFFF
	BoardWRAMBegin        = 0x02000000
	BoardWRAMEnd          = 0x0203FFFF
	Unused2Begin          = 0x02040000
	Unused2End            = 0x02FFFFFF
	ChipWRAMBegin         = 0x03000000
	ChipWRAMEnd           = 0x03007FFF
	Unused3Begin          = 0x03008000
	Unused3End            = 0x03FFFFFF
	IORegistersBegin      = 0x04000000
	IORegistersEnd        = 0x040003FE
	Unused4Begin          = 0x04000400
	Unused4End            = 0x04FFFFFF
	// Internal Display Memory
	PaletteRAMBegin = 0x05000000
	PaletteRAMEnd   = 0x050003FF
	Unused5Begin    = 0x05000400
	Unused5End      = 0x05FFFFFF
	VideoRAMBegin   = 0x06000000
	VideoRAMEmd     = 0x06017FFF
	Unused6Begin    = 0x06018000
	Unused6End      = 0x06FFFFFF
	OAMBegin        = 0x07000000
	OAMEnd          = 0x070003FF
	Unused7Begin    = 0x07000400
	Unused7End      = 0x07FFFFFF
	// External Memory (Game Pak)
	GPROM0Begin  = 0x08000000
	GPROM0End    = 0x09FFFFFF
	GPROM1Begin  = 0x0A000000
	GPROM1End    = 0x0BFFFFFF
	GPROM2Begin  = 0x0C000000
	GPROM2End    = 0x0DFFFFFF
	GPSRAMBegin  = 0x0E000000
	GPSRAMEnd    = 0x0E00FFFF
	Unused8Begin = 0x0E010000
	Unused8End   = 0x0FFFFFFF
	// Unused Memory Area
	Unused9Begin = 0x10000000
	Unused9End   = 0xFFFFFFFF
)
