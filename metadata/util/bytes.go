package util

import (
	"bytes"
	"unicode"
	"unicode/utf8"
)

// Clean removes non-graphic characters from the given string.
// Removable characters are the ones for which unicode.IsGraphic() returns false.
func PrintableBytes(data []byte) []byte {
	rs := bytes.Runes(data)
	var (
		size    int
		cleanRs = make([]rune, 0, len(rs))
	)
	for _, r := range rs {
		if !unicode.IsGraphic(r) {
			continue
		}
		cleanRs = append(cleanRs, r)
		size += utf8.RuneLen(r)
	}
	bs := make([]byte, size)
	var count int
	for _, r := range cleanRs {
		count += utf8.EncodeRune(bs[count:], r)
	}
	return bs
}
