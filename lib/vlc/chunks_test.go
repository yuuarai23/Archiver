package vlc

import (
	"reflect"
	"testing"
)

func TestBinaryChunks_ToHex(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinaryChunks
		want HexChunks
	}{
		{
			name: "test 1",
			bcs:  BinaryChunks{"00100000", "00110000", "00111100", "00011000", "01110111", "01001010", "11100100", "01001101", "00101000"},
			want: HexChunks{"20", "30", "3C", "18", "77", "4A", "E4", "4D", "28"},
		},
		{
			name: "test 2",
			bcs:  BinaryChunks{"00100000", "00110010", "00000000", "11100100", "01000001", "10010000", "00011101", "11010010", "01000010", "11110010", "01000101", "00100000", "10100000"},
			want: HexChunks{"20", "32", "00", "E4", "41", "90", "1D", "D2", "42", "F2", "45", "20", "A0"},
		},
		{
			name: "test 3",
			bcs:  BinaryChunks{"00100000", "00110010", "00000000", "11100100", "01000000", "10000110", "01000000", "01100100", "01011100", "10000100", "10010000", "10111001", "00010010", "01000101", "00100000", "10100000"},
			want: HexChunks{"20", "32", "00", "E4", "40", "86", "40", "64", "5C", "84", "90", "B9", "12", "45", "20", "A0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bcs.ToHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitByChunks(t *testing.T) {
	type args struct {
		bStr      string
		chunkSize int
	}

	tests := []struct {
		name string
		args args
		want BinaryChunks
	}{
		{
			name: "test 1",
			args: args{
				bStr:      "001000000011000000111100000110000111011101001010111001000100110100101",
				chunkSize: 8,
			},
			want: BinaryChunks{"00100000", "00110000", "00111100", "00011000", "01110111", "01001010", "11100100", "01001101", "00101000"},
		},
		{
			name: "test 2",
			args: args{
				bStr:      "001000000011001000000000111001000100000110010000000111011101001001000010111100100100010100100000101",
				chunkSize: 8,
			},
			want: BinaryChunks{"00100000", "00110010", "00000000", "11100100", "01000001", "10010000", "00011101", "11010010", "01000010", "11110010", "01000101", "00100000", "10100000"},
		},
		{
			name: "test 3",
			args: args{
				bStr:      "001000000011001000000000111001000100000010000110010000000110010001011100100001001001000010111001000100100100010100100000101",
				chunkSize: 8,
			},
			want: BinaryChunks{"00100000", "00110010", "00000000", "11100100", "01000000", "10000110", "01000000", "01100100", "01011100", "10000100", "10010000", "10111001", "00010010", "01000101", "00100000", "10100000"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitByChunks(tt.args.bStr, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHexChunks(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want HexChunks
	}{
		{
			name: "test 1",
			str:  "20 30 3C 18 77 4A E4 4D 28",
			want: HexChunks{"20", "30", "3C", "18", "77", "4A", "E4", "4D", "28"},
		},
		{
			name: "test 1",
			str:  "20 32 00 E4 41 90 1D D2 42 F2 45 20 A0",
			want: HexChunks{"20", "32", "00", "E4", "41", "90", "1D", "D2", "42", "F2", "45", "20", "A0"},
		},
		{
			name: "test 1",
			str:  "20 32 00 E4 40 86 40 64 5C 84 90 B9 12 45 20 A0",
			want: HexChunks{"20", "32", "00", "E4", "40", "86", "40", "64", "5C", "84", "90", "B9", "12", "45", "20", "A0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHexChunks(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHexChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexChunk_ToBinary(t *testing.T) {
	tests := []struct {
		name string
		hc   HexChunk
		want BinaryChunk
	}{
		{
			name: "test 1",
			hc:   HexChunk("2F"),
			want: BinaryChunk("00101111"),
		},
		{
			name: "test 2",
			hc:   HexChunk("00"),
			want: BinaryChunk("00000000"),
		},
		{
			name: "test 3",
			hc:   HexChunk("80"),
			want: BinaryChunk("10000000"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hc.ToBinary(); got != tt.want {
				t.Errorf("ToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexChunks_ToBinary(t *testing.T) {
	tests := []struct {
		name string
		hcs  HexChunks
		want BinaryChunks
	}{
		{
			name: "test 1",
			hcs:  HexChunks{"20", "30", "3C", "18", "77", "4A", "E4", "4D", "28"},
			want: BinaryChunks{"00100000", "00110000", "00111100", "00011000", "01110111", "01001010", "11100100", "01001101", "00101000"},
		},
		{
			name: "test 2",
			hcs:  HexChunks{"20", "32", "00", "E4", "41", "90", "1D", "D2", "42", "F2", "45", "20", "A0"},
			want: BinaryChunks{"00100000", "00110010", "00000000", "11100100", "01000001", "10010000", "00011101", "11010010", "01000010", "11110010", "01000101", "00100000", "10100000"},
		},
		{
			name: "test 3",
			hcs:  HexChunks{"20", "32", "00", "E4", "40", "86", "40", "64", "5C", "84", "90", "B9", "12", "45", "20", "A0"},
			want: BinaryChunks{"00100000", "00110010", "00000000", "11100100", "01000000", "10000110", "01000000", "01100100", "01011100", "10000100", "10010000", "10111001", "00010010", "01000101", "00100000", "10100000"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hcs.ToBinary(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexChunks_Join(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinaryChunks
		want string
	}{
		{
			name: "test 1",
			bcs:  BinaryChunks{"00100000", "00110000", "00111100", "00011000", "01110111", "01001010", "11100100", "01001101", "00101000"},
			want: "001000000011000000111100000110000111011101001010111001000100110100101000",
		},
		{
			name: "test 2",
			bcs:  BinaryChunks{"00100000", "00110010", "00000000", "11100100", "01000001", "10010000", "00011101", "11010010", "01000010", "11110010", "01000101", "00100000", "10100000"},
			want: "00100000001100100000000011100100010000011001000000011101110100100100001011110010010001010010000010100000",
		},
		{
			name: "test 3",
			bcs:  BinaryChunks{"00100000", "00110010", "00000000", "11100100", "01000000", "10000110", "01000000", "01100100", "01011100", "10000100", "10010000", "10111001", "00010010", "01000101", "00100000", "10100000"},
			want: "00100000001100100000000011100100010000001000011001000000011001000101110010000100100100001011100100010010010001010010000010100000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bcs.Join(); got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}
