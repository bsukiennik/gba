package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dumpCmd = &cobra.Command{
	Use:     "dump",
	Aliases: []string{"display"},
	Short:   "Dump the header of the gba rom",
	Long:    "Dump the header of the gba rom",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		err, res := dump(args[0])
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", res)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(dumpCmd)
}

func dumpRomEntryPoint(romEntryPoint [4]byte) (result string) {
	result += " ├── entry:\n"
	result += fmt.Sprintf(" │   └── raw: 0x%x\n", romEntryPoint)
	return result
}

func dumpNintendoLogo(nintendoLogo [156]byte) (result string) {
	result += " ├── nintendo logo:\n"
	validNintendoLogo := [156]byte{
		0x24, 0xFF, 0xAE, 0x51, 0x69, 0x9A, 0xA2, 0x21, 0x3D, 0x84, 0x82, 0x0A, 0x84, 0xE4, 0x09, 0xAD,
		0x11, 0x24, 0x8B, 0x98, 0xC0, 0x81, 0x7F, 0x21, 0xA3, 0x52, 0xBE, 0x19, 0x93, 0x09, 0xCE, 0x20,
		0x10, 0x46, 0x4A, 0x4A, 0xF8, 0x27, 0x31, 0xEC, 0x58, 0xC7, 0xE8, 0x33, 0x82, 0xE3, 0xCE, 0xBF,
		0x85, 0xF4, 0xDF, 0x94, 0xCE, 0x4B, 0x09, 0xC1, 0x94, 0x56, 0x8A, 0xC0, 0x13, 0x72, 0xA7, 0xFC,
		0x9F, 0x84, 0x4D, 0x73, 0xA3, 0xCA, 0x9A, 0x61, 0x58, 0x97, 0xA3, 0x27, 0xFC, 0x03, 0x98, 0x76,
		0x23, 0x1D, 0xC7, 0x61, 0x03, 0x04, 0xAE, 0x56, 0xBF, 0x38, 0x84, 0x00, 0x40, 0xA7, 0x0E, 0xFD,
		0xFF, 0x52, 0xFE, 0x03, 0x6F, 0x95, 0x30, 0xF1, 0x97, 0xFB, 0xC0, 0x85, 0x60, 0xD6, 0x80, 0x25,
		0xA9, 0x63, 0xBE, 0x03, 0x01, 0x4E, 0x38, 0xE2, 0xF9, 0xA2, 0x34, 0xFF, 0xBB, 0x3E, 0x03, 0x44,
		0x78, 0x00, 0x90, 0xCB, 0x88, 0x11, 0x3A, 0x94, 0x65, 0xC0, 0x7C, 0x63, 0x87, 0xF0, 0x3C, 0xAF,
		0xD6, 0x25, 0xE4, 0x8B, 0x38, 0x0A, 0xAC, 0x72, 0x21, 0xD4, 0xF8, 0x07,
	}

	var status string
	if nintendoLogo == validNintendoLogo {
		status = "valid"
	} else {
		status = "not valid"
	}

	var debugging string
	validNintendoLogo[152] = 0xA5
	if nintendoLogo == validNintendoLogo {
		debugging = "enabled"
	} else {
		debugging = "disabled"
	}

	result += fmt.Sprintf(" │   ├── status: %s\n", status)
	result += fmt.Sprintf(" │   └── debugging: %s\n", debugging)
	return result
}

func dumpGameTitle(gameTitle [12]byte) (result string) {
	return fmt.Sprintf(" ├── game title: %s\n", gameTitle)
}

func dumpGameCode(gameCode [4]byte) (result string) {
	result += " ├── game code:\n"
	result += fmt.Sprintf(" │   ├── code: %s\n", gameCode)

	var language string
	switch gameCode[3] {
	case 'J':
		language = "Japan"
	case 'P':
		language = "Europe/Elsewhere"
	case 'F':
		language = "French"
	case 'S':
		language = "Spanish"
	case 'E':
		language = "USA/English"
	case 'D':
		language = "German"
	case 'I':
		language = "Italian"
	default:
		language = "undefined"
	}
	result += fmt.Sprintf(" │   └── language: %s\n", language)
	return result
}

func dumpMakerCode(makerCode [2]byte) (result string) {
	result += " ├── maker code:\n"
	result += fmt.Sprintf(" │   ├── id: %s\n", makerCode)

	var developer string
	switch makerCode {
	case [2]byte{'0', '1'}:
		developer = "nintendo"
	default:
		developer = "unknown developer"
	}
	result += fmt.Sprintf(" │   └── developer: %s\n", developer)
	return result
}

func dumpFixedValue(fixedValue [1]byte) (result string) {
	var valid string
	if fixedValue[0] == 0x96 {
		valid = "valid"
	} else {
		valid = "invalid"
	}
	result += fmt.Sprintf(" ├── fixed value: %s (%.2xh)\n", valid, fixedValue[0])
	return result
}

func dumpMainUnit(mainUnit [1]byte) (result string) {
	var valid string
	if mainUnit[0] == 0x00 {
		valid = ""
	} else {
		valid = " (should be 00h)"
	}

	result += fmt.Sprintf(" ├── unit code: %.2xh%x\n", mainUnit[0], valid)
	return result
}

func dumpDeviceType(deviceType [1]byte) (result string) {
	var valid string
	if deviceType[0] == 0x00 {
		valid = ""
	} else {
		valid = " (should be 00h)"
	}

	result += fmt.Sprintf(" ├── device type: %.2xh%x\n", deviceType[0], valid)
	return result
}

func dumpReservedArea(reservedAreaOne [7]byte, reservedAreaTwo [2]byte) (result string) {
	var error string = ""
	for _, value := range reservedAreaOne {
		if value != 0x00 {
			error += "(invalid first area)"
			break
		}
	}
	for _, value := range reservedAreaTwo {
		if value != 0x00 {
			error += "(invalid second area)"
			break
		}
	}

	var valid string
	if error == "" {
		valid = "valid"
	} else {
		valid = "invalid " + error
	}
	result += fmt.Sprintf(" ├── reserved area: %s\n", valid)
	return result
}

func dumpSoftwareVersion(softwareVersion [1]byte) (result string) {
	var usually string
	if softwareVersion[0] == 0x00 {
		usually = ""
	} else {
		usually = " (usally 00h)"
	}

	result += fmt.Sprintf(" ├── software version: %.2xh%s\n", softwareVersion[0], usually)
	return result
}

func dumpComplementCheck(complementCheck [1]byte, rom []byte) (result string) {
	result += " └── check sum:\n"

	var checkSum int

	for i := 0x0A0; i <= 0x0BC; i++ {
		checkSum -= int(rom[i])
	}
	checkSum = (checkSum - 0x19) & 0xFF

	var valid bool
	if checkSum == int(complementCheck[0]) {
		valid = true
	} else {
		valid = false
	}

	result += fmt.Sprintf("     ├── valid: %t\n", valid)
	result += fmt.Sprintf("     ├── rom: %.2x\n", complementCheck)
	result += fmt.Sprintf("     └── our: %.2x\n", checkSum)
	return result
}
