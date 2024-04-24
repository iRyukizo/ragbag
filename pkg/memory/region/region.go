package region

// Those const represent the GBA General Internal Memory Map.
const (
	BIOSBegin        uint = 0x0000_0000
	BIOSEnd               = 0x0000_3FFF
	Unused1Begin          = 0x0000_4000
	Unused1End            = 0x01FF_FFFF
	BoardWRAMBegin        = 0x0200_0000
	BoardWRAMEnd          = 0x0203_FFFF
	Unused2Begin          = 0x0204_0000
	Unused2End            = 0x02FF_FFFF
	ChipWRAMBegin         = 0x0300_0000
	ChipWRAMEnd           = 0x0300_7FFF
	Unused3Begin          = 0x0300_8000
	Unused3End            = 0x03FF_FFFF
	IORegistersBegin      = 0x0400_0000
	IORegistersEnd        = 0x0400_03FE
	Unused4Begin          = 0x0400_0400
	Unused4End            = 0x04FF_FFFF
)

// Those const represent the GBA Internal Display Memory Map.
const (
	PaletteRAMBegin uint = 0x0500_0000
	PaletteRAMEnd        = 0x0500_03FF
	Unused5Begin         = 0x0500_0400
	Unused5End           = 0x05FF_FFFF
	VideoRAMBegin        = 0x0600_0000
	VideoRAMEmd          = 0x0601_7FFF
	Unused6Begin         = 0x0601_8000
	Unused6End           = 0x06FF_FFFF
	OAMBegin             = 0x0700_0000
	OAMEnd               = 0x0700_03FF
	Unused7Begin         = 0x0700_0400
	Unused7End           = 0x07FF_FFFF
)

// Those const represent the GBA External Memory (Game Pak) Map.
const (
	GPROM0Begin  uint = 0x0800_0000
	GPROM0End         = 0x09FF_FFFF
	GPROM1Begin       = 0x0A00_0000
	GPROM1End         = 0x0BFF_FFFF
	GPROM2Begin       = 0x0C00_0000
	GPROM2End         = 0x0DFF_FFFF
	GPSRAMBegin       = 0x0E00_0000
	GPSRAMEnd         = 0x0E00_FFFF
	Unused8Begin      = 0x0E01_0000
	Unused8End        = 0x0FFF_FFFF
)

// Those const represent the GBA Unused Memory Map.
const (
	Unused9Begin uint = 0x1000_0000
	Unused9End        = 0xFFFF_FFFF
)
