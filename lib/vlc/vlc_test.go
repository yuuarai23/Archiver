package vlc

import (
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
				t.Errorf("prepareText() = %v, want %v", got, tt.want)
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
				t.Errorf("encodeBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "test 1",
			str:  "My name is Ted",
			want: "20 30 3C 18 77 4A E4 4D 28",
		},
		{
			name: "test 2",
			str:  "MY NaMe iS tED",
			want: "20 32 00 E4 41 90 1D D2 42 F2 45 20 A0",
		},
		{
			name: "test 3",
			str:  "MY NAME IS TED",
			want: "20 32 00 E4 40 86 40 64 5C 84 90 B9 12 45 20 A0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.str); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
