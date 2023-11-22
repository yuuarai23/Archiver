package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type BinaryChunk string
type HexChunk string
type HexChunks []HexChunk
type BinaryChunks []BinaryChunk
type encodingTable map[rune]string

const chunksSize = 8

func Encode(str string) string {
	str = prepareText(str)

	bStr := encodeBin(str)

	chunks := splitByChunks(bStr, chunksSize)

	fmt.Println(chunks)

	// bytes to hex -> '20 30 3C'
	chunks.ToHex()

	// return hexChunksStr

	return ""
}

func (bcs BinaryChunks) ToHex() HexChunks {
	res := make(HexChunks, 0, len(bcs))

	for _, chunk := range bcs {
		hexChunk := chunk.ToHex()

		res = append(res, hexChunk)
	}

	return res
}

func (bc BinaryChunk) ToHex() HexChunk {
	num, err := strconv.ParseInt(string(bc), 2, chunksSize)
	if err != nil {
		panic("can't parse binary chunk: " + err.Error())
	}

	res := strings.ToUpper(fmt.Sprintf("%x", num))

	if len(res) == 1 {
		res = "0" + res
	}

	return HexChunk(res)
}

// splitByChunks splits binary string by chunks with given size,
// i.g.: '100101011001010110010101' -> '10010101 10010101 10010101'
func splitByChunks(bStr string, chunkSize int) BinaryChunks {
	strLen := utf8.RuneCountInString(bStr)

	chunksCount := strLen / chunkSize

	if strLen/chunkSize != 0 {
		chunksCount++
	}

	res := make(BinaryChunks, 0, chunksCount)

	var buf strings.Builder

	for i, ch := range bStr {
		buf.WriteString(string(ch))

		if (i+1)%chunkSize == 0 {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}

	if buf.Len() != 0 {
		lastChunk := buf.String()
		lastChunk += strings.Repeat("0", chunkSize-len(lastChunk))

		res = append(res, BinaryChunk(lastChunk))
	}

	return res
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
