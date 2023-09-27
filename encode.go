package uuid7

import (
	"errors"
	"strings"

	"github.com/mr-tron/base58"
)

func NewString() string {
	return EncodeBase58(MustUUID7())
}

func encodeBase58Raw(u UUID) string {
	return FastBase58EncodingAlphabet(u[:], BTCAlphabet)
}

func EncodeBase58(u UUID) string {
	s := FastBase58EncodingAlphabet(u[:], BTCAlphabet)
	if len(s) != 22 {
		s = strings.Repeat("1", 22-len(s)) + s // pad with leading "zeroes" (1 in BTC base58)
	}
	return s[0:9] + "_" + s[9:]
}

func DecodeBase58(s string) (UUID, error) {
	if len(s) != 23 {
		return UUID{}, errors.New("uuid7 base58: invalid length")
	}
	if s[9] != '_' {
		return UUID{}, errors.New("uuid7 base58: invalid separator")
	}
	s = s[0:9] + s[10:]
	d, err := base58.Decode(s)
	if err != nil {
		return UUID{}, err
	}
	if len(d) > 16 {
		d = d[len(d)-16:] // remove leading "zeroes" (1 in BTC base58)
	}

	var result UUID
	copy(result[:], d)
	return result, nil
}
