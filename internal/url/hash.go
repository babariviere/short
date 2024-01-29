package url

import (
	"hash/fnv"
	"strings"
)

// GenerateHash hash the `url` param into a 8 characters base62 using fnv algorithm.
func GenerateHash(url string) (string, error) {
	hasher := fnv.New32a()
	_, err := hasher.Write([]byte(url))
	if err != nil {
		return "", err
	}
	return encodeToBase62(hasher.Sum32()), nil
}

var base62 = [62]rune{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
}

func encodeToBase62(hash uint32) string {
	var sb strings.Builder
	sb.Grow(8)

	for hash > 0 {
		sb.WriteRune(base62[hash%62])
		hash /= 62
	}

	return sb.String()
}
