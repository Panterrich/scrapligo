package channel

import (
	"fmt"
	"unicode"
)

// function for detailed analyze bytes
func analyzeBytes(data []byte) {
	fmt.Printf("Raw bytes (hex): ")
	for i, b := range data {
		fmt.Printf("%02x ", b)
		if (i+1)%16 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
	fmt.Println()

	fmt.Printf("Detailed analysis:\n")
	fmt.Printf("Index | Hex  | Dec  | Char  | Unicode | Description\n")
	fmt.Printf("------|------|------|-------|---------|------------\n")

	for i, b := range data {
		r := rune(b)

		var description string
		var charRep string

		switch {
		case b == 0:
			description = "NULL"
			charRep = "\\0"
		case b == 7:
			description = "BEL (bell)"
			charRep = "\\a"
		case b == 8:
			description = "BS (backspace)"
			charRep = "\\b"
		case b == 9:
			description = "TAB"
			charRep = "\\t"
		case b == 10:
			description = "LF (newline)"
			charRep = "\\n"
		case b == 11:
			description = "VT (vertical tab)"
			charRep = "\\v"
		case b == 12:
			description = "FF (form feed)"
			charRep = "\\f"
		case b == 13:
			description = "CR (carriage return)"
			charRep = "\\r"
		case b == 27:
			description = "ESC (escape)"
			charRep = "\\e"
		case b == 32:
			description = "Space"
			charRep = "␣"
		case b == 127:
			description = "DEL"
			charRep = "\\x7f"
		case unicode.IsPrint(r) && b >= 32 && b <= 126:
			description = "Printable ASCII"
			charRep = string(b)
		case b >= 128:
			description = "Extended ASCII/UTF-8"
			charRep = fmt.Sprintf("\\x%02x", b)
		default:
			description = "Control character"
			charRep = fmt.Sprintf("\\x%02x", b)
		}

		unicodeRep := fmt.Sprintf("U+%04X", r)

		fmt.Printf("%5d| 0x%02x | %3d | %5s | %7s | %s\n",
			i, b, b, charRep, unicodeRep, description)
	}
}

// function for comparing string and byte presentation
func CompareOutputs(data []byte) {
	fmt.Printf("Regular string output (%%s): \"%s\"\n", string(data))
	fmt.Printf("Regular string output (%%q): %q\n", string(data))
	fmt.Printf("Regular byte output (%%v): %v\n", data)
	fmt.Println()

	analyzeBytes(data)
}
