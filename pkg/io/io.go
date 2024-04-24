package io

// Those const represent GBA LCD I/O Registers Map.
const (
	DISPCNT  uint = 0x0400_0000 //  2    R/W  LCD Control
	Unused0       = 0x0400_0002 //  2    R/W  Undocumented - Green Swap
	DISPSTAT      = 0x0400_0004 //  2    R/W  General LCD Status (STAT,LYC)
	VCOUNT        = 0x0400_0006 //  2    R    Vertical Counter (LY)
	BG0CNT        = 0x0400_0008 //  2    R/W  BG0 Control
	BG1CNT        = 0x0400_000A //  2    R/W  BG1 Control
	BG2CNT        = 0x0400_000C //  2    R/W  BG2 Control
	BG3CNT        = 0x0400_000E //  2    R/W  BG3 Control
	BG0HOFS       = 0x0400_0010 //  2    W    BG0 X-Offset
	BG0VOFS       = 0x0400_0012 //  2    W    BG0 Y-Offset
	BG1HOFS       = 0x0400_0014 //  2    W    BG1 X-Offset
	BG1VOFS       = 0x0400_0016 //  2    W    BG1 Y-Offset
	BG2HOFS       = 0x0400_0018 //  2    W    BG2 X-Offset
	BG2VOFS       = 0x0400_001A //  2    W    BG2 Y-Offset
	BG3HOFS       = 0x0400_001C //  2    W    BG3 X-Offset
	BG3VOFS       = 0x0400_001E //  2    W    BG3 Y-Offset
	BG2PA         = 0x0400_0020 //  2    W    BG2 Rotation/Scaling Parameter A (dx)
	BG2PB         = 0x0400_0022 //  2    W    BG2 Rotation/Scaling Parameter B (dmx)
	BG2PC         = 0x0400_0024 //  2    W    BG2 Rotation/Scaling Parameter C (dy)
	BG2PD         = 0x0400_0026 //  2    W    BG2 Rotation/Scaling Parameter D (dmy)
	BG2X          = 0x0400_0028 //  4    W    BG2 Reference Point X-Coordinate
	BG2Y          = 0x0400_002C //  4    W    BG2 Reference Point Y-Coordinate
	BG3PA         = 0x0400_0030 //  2    W    BG3 Rotation/Scaling Parameter A (dx)
	BG3PB         = 0x0400_0032 //  2    W    BG3 Rotation/Scaling Parameter B (dmx)
	BG3PC         = 0x0400_0034 //  2    W    BG3 Rotation/Scaling Parameter C (dy)
	BG3PD         = 0x0400_0036 //  2    W    BG3 Rotation/Scaling Parameter D (dmy)
	BG3X          = 0x0400_0038 //  4    W    BG3 Reference Point X-Coordinate
	BG3Y          = 0x0400_003C //  4    W    BG3 Reference Point Y-Coordinate
	WIN0H         = 0x0400_0040 //  2    W    Window 0 Horizontal Dimensions
	WIN1H         = 0x0400_0042 //  2    W    Window 1 Horizontal Dimensions
	WIN0V         = 0x0400_0044 //  2    W    Window 0 Vertical Dimensions
	WIN1V         = 0x0400_0046 //  2    W    Window 1 Vertical Dimensions
	WININ         = 0x0400_0048 //  2    R/W  Inside of Window 0 and 1
	WINOUT        = 0x0400_004A //  2    R/W  Inside of OBJ Window & Outside of Windows
	MOSAIC        = 0x0400_004C //  2    W    Mosaic Size
	Unused1       = 0x0400_004E //       -    Not used
	BLDCNT        = 0x0400_0050 //  2    R/W  Color Special Effects Selection
	BLDALPHA      = 0x0400_0052 //  2    R/W  Alpha Blending Coefficients
	BLDY          = 0x0400_0054 //  2    W    Brightness (Fade-In/Out) Coefficient
	Unused2       = 0x0400_0056 //       -    Not used
)

// Those const represent GBA Sounds Registers Map.
const (
	SOUND1CNT_L uint = 0x0400_0060 //  2    R/W  SOChannel 1 Sweep register       (NR10)
	SOUND1CNT_H      = 0x0400_0062 //  2    R/W  SOChannel 1 Duty/Length/Envelope (NR11, NR12)
	SOUND1CNT_X      = 0x0400_0064 //  2    R/W  SOChannel 1 Frequency/Control    (NR13, NR14)
	Unused3          = 0x0400_0066 //       -    Not used
	SOUND2CNT_L      = 0x0400_0068 //  2    R/W  SOChannel 2 Duty/Length/Envelope (NR21, NR22)
	Unused4          = 0x0400_006A //       -    Not used
	SOUND2CNT_H      = 0x0400_006C //  2    R/W  SOChannel 2 Frequency/Control    (NR23, NR24)
	Unused5          = 0x0400_006E //       -    Not used
	SOUND3CNT_L      = 0x0400_0070 //  2    R/W  SOChannel 3 Stop/Wave RAM select (NR30)
	SOUND3CNT_H      = 0x0400_0072 //  2    R/W  SOChannel 3 Length/Volume        (NR31, NR32)
	SOUND3CNT_X      = 0x0400_0074 //  2    R/W  SOChannel 3 Frequency/Control    (NR33, NR34)
	Unused6          = 0x0400_0076 //       -    Not used
	SOUND4CNT_L      = 0x0400_0078 //  2    R/W  SOChannel 4 Length/Envelope      (NR41, NR42)
	Unused7          = 0x0400_007A //       -    Not used
	SOUND4CNT_H      = 0x0400_007C //  2    R/W  SOChannel 4 Frequency/Control    (NR43, NR44)
	Unused8          = 0x0400_007E //       -    Not used
	SOUNDCNT_L       = 0x0400_0080 //  2    R/W  SOControl Stereo/Volume/Enable   (NR50, NR51)
	SOUNDCNT_H       = 0x0400_0082 //  2    R/W  SOControl Mixing/DMA Control
	SOUNDCNT_X       = 0x0400_0084 //  2    R/W  SOControl Sound on/off           (NR52)
	Unused9          = 0x0400_0086 //       -    Not used
	SOUNDBIAS        = 0x0400_0088 //  2    BIOS SOSound PWM Control
	Unused10         = 0x0400_008A //  ..   -    Not used
	WAVE_RAM         = 0x0400_0090 // 2x10h R/W  Channel 3 Wave Pattern RAM (2 banks!!)
	FIFO_A           = 0x0400_00A0 //  4    W    Channel A FIFO, Data 0-3
	FIFO_B           = 0x0400_00A4 //  4    W    Channel B FIFO, Data 0-3
	Unused11         = 0x0400_00A8 //       -    Not used
)

// Those const represent GBA DMA Transfer Channels Map.
const (
	DMA0SAD   uint = 0x0400_00B0 //  4    W    DMA 0 Source Address
	DMA0DAD        = 0x0400_00B4 //  4    W    DMA 0 Destination Address
	DMA0CNT_L      = 0x0400_00B8 //  2    W    DMA 0 Word Count
	DMA0CNT_H      = 0x0400_00BA //  2    R/W  DMA 0 Control
	DMA1SAD        = 0x0400_00BC //  4    W    DMA 1 Source Address
	DMA1DAD        = 0x0400_00C0 //  4    W    DMA 1 Destination Address
	DMA1CNT_L      = 0x0400_00C4 //  2    W    DMA 1 Word Count
	DMA1CNT_H      = 0x0400_00C6 //  2    R/W  DMA 1 Control
	DMA2SAD        = 0x0400_00C8 //  4    W    DMA 2 Source Address
	DMA2DAD        = 0x0400_00CC //  4    W    DMA 2 Destination Address
	DMA2CNT_L      = 0x0400_00D0 //  2    W    DMA 2 Word Count
	DMA2CNT_H      = 0x0400_00D2 //  2    R/W  DMA 2 Control
	DMA3SAD        = 0x0400_00D4 //  4    W    DMA 3 Source Address
	DMA3DAD        = 0x0400_00D8 //  4    W    DMA 3 Destination Address
	DMA3CNT_L      = 0x0400_00DC //  2    W    DMA 3 Word Count
	DMA3CNT_H      = 0x0400_00DE //  2    R/W  DMA 3 Control
	Unused12       = 0x0400_00E0 //       -    Not used
)

// Those const represent GBA Timer Registers Map.
const (
	TM0CNT_L uint = 0x0400_0100 //  2    R/W  Timer 0 Counter/Reload
	TM0CNT_H      = 0x0400_0102 //  2    R/W  Timer 0 Control
	TM1CNT_L      = 0x0400_0104 //  2    R/W  Timer 1 Counter/Reload
	TM1CNT_H      = 0x0400_0106 //  2    R/W  Timer 1 Control
	TM2CNT_L      = 0x0400_0108 //  2    R/W  Timer 2 Counter/Reload
	TM2CNT_H      = 0x0400_010A //  2    R/W  Timer 2 Control
	TM3CNT_L      = 0x0400_010C //  2    R/W  Timer 3 Counter/Reload
	TM3CNT_H      = 0x0400_010E //  2    R/W  Timer 3 Control
	Unused13      = 0x0400_0110 //       -    Not used
)

// Those const represent GBA Serial Communication (1) Map.
const (
	SIODATA32   uint = 0x0400_0120 //  4    R/W  SIO Data (Normal-32bit Mode; shared with below)
	SIOMULTI0        = 0x0400_0120 //  2    R/W  SIO Data 0 (Parent)    (Multi-Player Mode)
	SIOMULTI1        = 0x0400_0122 //  2    R/W  SIO Data 1 (1st Child) (Multi-Player Mode)
	SIOMULTI2        = 0x0400_0124 //  2    R/W  SIO Data 2 (2nd Child) (Multi-Player Mode)
	SIOMULTI3        = 0x0400_0126 //  2    R/W  SIO Data 3 (3rd Child) (Multi-Player Mode)
	SIOCNT           = 0x0400_0128 //  2    R/W  SIO Control Register
	SIOMLT_SEND      = 0x0400_012A //  2    R/W  SIO Data (Local of MultiPlayer; shared below)
	SIODATA8         = 0x0400_012A //  2    R/W  SIO Data (Normal-8bit and UART Mode)
	Unused14         = 0x0400_012C //       -    Not used
)

// Those const represent GBA Keypad Input Map.
const (
	KEYINPUT uint = 0x0400_0130 //  2    R    Key Status
	KEYCNT        = 0x0400_0132 //  2    R/W  Key Interrupt Control
)

// Those const represent GBA Serial Communication (2) Map.
const (
	RCNT      = 0x0400_0134 //  2    R/W  SIO Mode Select/General Purpose Data
	IR        = 0x0400_0136 //  -    -    Ancient - Infrared Register (Prototypes only)
	Unused15  = 0x0400_0138 //       -    Not used
	JOYCNT    = 0x0400_0140 //  2    R/W  SIO JOY Bus Control
	Unused16  = 0x0400_0142 //       -    Not used
	JOY_RECV  = 0x0400_0150 //  4    R/W  SIO JOY Bus Receive Data
	JOY_TRANS = 0x0400_0154 //  4    R/W  SIO JOY Bus Transmit Data
	JOYSTAT   = 0x0400_0158 //  2    R/?  SIO JOY Bus Receive Status
	Unused17  = 0x0400_015A //       -    Not used
)

// Those const represent GBA Interrupt, Waitstate, and Power-Down Control Map.
const (
	IE       uint = 0x0400_0200 //  2    R/W  Interrupt Enable Register
	IF            = 0x0400_0202 //  2    R/W  Interrupt Request Flags / IRQ Acknowledge
	WAITCNT       = 0x0400_0204 //  2    R/W  Game Pak Waitstate Control
	Unused18      = 0x0400_0206 //       -    Not used
	IME           = 0x0400_0208 //  2    R/W  Interrupt Master Enable Register
	Unused19      = 0x0400_020A //       -    Not used
	POSTFLG       = 0x0400_0300 //  1    R/W  Undocumented - Post Boot Flag
	HALTCNT       = 0x0400_0301 //  1    W    Undocumented - Power Down Control
	Unused20      = 0x0400_0302 //       -    Not used
	Unused21      = 0x0400_0410 //  ?    ?    Undocumented - Purpose Unknown / Bug ??? 0FFh
	Unused22      = 0x0400_0411 //       -    Not used
	Unused23      = 0x0400_0800 //  4    R/W  Undocumented - Internal Memory Control (R/W)
	Unused24      = 0x0400_0804 //       -    Not used
	Unused25      = 0x0400_0800 //  4    R/W  Mirrors of 4000800h (repeated each 64K)
	_3DS          = 0x0470_0000 //  4    W    Disable ARM7 bootrom overlay (3DS only)
)
