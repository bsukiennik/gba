package cmd

import (
	"encoding/binary"
	"fmt"
	"os"
)

//   000h    4     ROM Entry Point  (32bit ARM branch opcode, eg. "B rom_start")
//   004h    156   Nintendo Logo    (compressed bitmap, required!)
//   0A0h    12    Game Title       (uppercase ascii, max 12 characters)
//   0ACh    4     Game Code        (uppercase ascii, 4 characters)
//   0B0h    2     Maker Code       (uppercase ascii, 2 characters)
//   0B2h    1     Fixed value      (must be 96h, required!)
//   0B3h    1     Main unit code   (00h for current GBA models)
//   0B4h    1     Device type      (usually 00h) (bit7=DACS/debug related)
//   0B5h    7     Reserved Area    (should be zero filled)
//   0BCh    1     Software version (usually 00h)
//   0BDh    1     Complement check (header checksum, required!)
//   0BEh    2     Reserved Area    (should be zero filled)

type Header struct {
	RomEntryPoint   [4]byte
	NintendoLogo    [156]byte
	GameTitle       [12]byte
	GameCode        [4]byte
	MakerCode       [2]byte
	FixedValue      [1]byte
	MainUnit        [1]byte
	DeviceType      [1]byte
	ReservedAreaOne [7]byte
	SoftwareVersion [1]byte
	ComplementCheck [1]byte
	ReservedAreaTwo [2]byte
}

func dump(rom string) (e error, result string) {
	file, err := os.Open(rom)
	if err != nil {
		return err, ""
	}
	defer file.Close()

	var header Header
	err = binary.Read(file, binary.LittleEndian, &header)
	if err != nil {
		return err, ""

	}

	fmt.Printf("%s\n", header.GameTitle)
	return nil, "pass"
}
