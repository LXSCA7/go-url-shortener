package base62

import "strings"

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	base     = uint64(len(alphabet))
)

func Encode(id int64) string {
	if id == 0 {
		return string(alphabet[0])
	}

	num := uint64(id)
	var sb strings.Builder
	for num > 0 {
		remainder := num % base
		sb.WriteByte(alphabet[remainder])
		num /= base
	}

	return reverse(sb.String())
}

func Decode(token string) int64 {
	var num uint64

	for i := 0; i < len(token); i++ {
		char := token[i]
		index := strings.IndexByte(alphabet, char)
		num = num*base + uint64(index)
	}

	return int64(num)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
