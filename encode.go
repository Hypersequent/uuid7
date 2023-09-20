package uuid7

import (
	"errors"
	"github.com/mr-tron/base58"
)

func NewString() string {
	return EncodeBase58(MustUUID7())
}

func encodeBase58Raw(u UUID) string {
	return base58.Encode(u[:])
}

func EncodeBase58(u UUID) string {
	s := base58.Encode(u[:])
	if len(s) != 22 {
		return "1" + s // pad with leading "zeroes" (1 in BTC base58)
	}
	return s
}

func DecodeBase58(s string) (UUID, error) {
	d, err := base58.Decode(s)
	if err != nil {
		return UUID{}, err
	}
	if len(d) == 17 {
		d = d[1:] // remove leading "zeroes" (1 in BTC base58)
	}
	if len(d) != 16 {
		return UUID{}, errors.New("invalid base58 UUID length")
	}
	return UUID(d), nil
}
