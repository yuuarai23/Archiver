package vlc

import (
	"strings"
	"unicode"
)

func Encode(str string) string {
	str = prepareText(str)

	chunks := splitByChunks(encodeBin(str), chunksSize)

	return chunks.ToHex().ToString()
}

func Decode(encodedText string) string {
	// hChunks := NewHexChunks(encodedText)

	// bChunks := hChunks.ToBinary()

	return ""
}

// encodeBin encodes str into binary codes string without spaces.
func encodeBin(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		buf.WriteString(bin(ch))
	}

	return buf.String()
}

func bin(ch rune) string {
	table := getEncodingTable()

	res, ok := table[ch]
	if !ok {
		panic("unknown character: " + string(ch))
	}
	return res
}

func getEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		'a': "011",
		'b': "0000010",
		'c': "000101",
		'd': "00101",
		'e': "101",
		'f': "000100",
		'g': "0000100",
		'h': "0011",
		'i': "01001",
		'j': "000000001",
		'k': "0000000001",
		'l': "001001",
		'm': "000011",
		'n': "10000",
		'o': "10001",
		'p': "0000101",
		'q': "000000000001",
		'r': "01000",
		's': "0101",
		't': "1001",
		'u': "00011",
		'v': "00000001",
		'w': "0000011",
		'x': "00000000001",
		'y': "0000001",
		'z': "000000000000",
		'!': "001000",
	}
}

// prepareText prepares text to be fit for encode:
// changes upper case letters to: ! + lower case letter
// i.g.: My name is Ted -> !my name is !ted
func prepareText(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		if unicode.IsUpper(ch) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(ch))
		} else {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}
