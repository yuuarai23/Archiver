package vlc

import (
	"reflect"
	"testing"
)

func Test_prepareText(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "test 1",
			str:  "My name is Ted",
			want: "!my name is !ted",
		},
		{
			name: "test 2",
			str:  "MY NaMe iS tED",
			want: "!m!y !na!me i!s t!e!d",
		},
		{
			name: "test 3",
			str:  "MY NAME IS TED",
			want: "!m!y !n!a!m!e !i!s !t!e!d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareText(tt.str); got != tt.want {
				t.Errorf("prepareText() = {%v}, want {%v}", got, tt.want)
			}
		})
	}
}

func Test_encodeBin(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "test 1",
			str:  "!my name is !ted",
			want: "001000000011000000111100000110000111011101001010111001000100110100101",
		},
		{
			name: "test 2",
			str:  "!m!y !na!me i!s t!e!d",
			want: "001000000011001000000000111001000100000110010000000111011101001001000010111100100100010100100000101",
		},
		{
			name: "test 3",
			str:  "!m!y !n!a!m!e !i!s !t!e!d",
			want: "001000000011001000000000111001000100000010000110010000000110010001011100100001001001000010111001000100100100010100100000101",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeBin(tt.str); got != tt.want {
				t.Errorf("encodeBin() = {%v}, want {%v}", got, tt.want)
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
				t.Errorf("splitByChunks() = {%v}, want {%v}", got, tt.want)
			}
		})
	}
}
