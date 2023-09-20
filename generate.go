package uuid7

import (
	"crypto/rand"
	"encoding/binary"
	"math"
	"math/big"
	"time"
)

type UUID [16]byte

func MustUUID7() UUID {
	uuid, err := UUID7()
	if err != nil {
		panic(err)
	}
	return uuid
}

func UUID7() (UUID, error) {
	return FromTime(time.Now())
}

func FromTime(uuidTime time.Time) (UUID, error) {
	milliseconds := uint64(uuidTime.UnixNano() / int64(time.Millisecond))
	nanoseconds := uint64(uuidTime.UnixNano() % int64(time.Millisecond))

	// Calculate the 12-bit sub-millisecond precision time
	subMillisecond := float64(nanoseconds) / float64(time.Millisecond)
	precisionBitsValue := uint16(math.Floor(subMillisecond * 4096))

	// Initialize a buffer to hold the big-endian bytes
	var uuidBytes UUID

	// Version field set to 0b0111 (7)
	version := uint16(7)

	// Manually construct the first 64-bit field
	// 48 bits for the timestamp, 4 bits for version, and 12 bits for sub-ms time
	sixtyFourBitField := (milliseconds << 16) | ((uint64(version) << 12) & 0xF000) | uint64(precisionBitsValue)

	binary.BigEndian.PutUint64(uuidBytes[0:], sixtyFourBitField)

	// Generate a 64-bit random number
	randomValue, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		return UUID{}, err
	}

	// Clear the top two bits and set them to 0b10
	random64 := randomValue.Uint64()
	random64 &= math.MaxUint64 >> 2 // Clear top two bits
	random64 |= uint64(2) << 62     // Set top bits to 0b10

	binary.BigEndian.PutUint64(uuidBytes[8:], random64)

	return uuidBytes, nil
}
