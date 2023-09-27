package uuid7

import (
	"fmt"
	"testing"
	"time"
)

func TestOutput(t *testing.T) {
	var uuids [10]string
	for i := 0; i < 10; i++ {
		uuids[i] = NewString()
		//time.Sleep(1 * time.Millisecond)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(uuids[i])
	}
}

func TestZero(t *testing.T) {
	uuid := UUID{}
	encoded := EncodeBase58(uuid)
	if encoded != "111111111_1111111111111" {
		t.Errorf("encoded string should be 111111111_1111111111111, got %s", encoded)
	}
}

func TestOne(t *testing.T) {
	testVecs := []string{
		"18AQGAut7_N92awznwCnjuR",
		"112d7dWtQ_Mvj9WttA3mMnX",
		"1111NKioe_UVktgzXLJ1B3u",
		"111115qCH_TcgbQwpvYZQ9d",
		"11111126U_w2Vvq8EnJ7hRH",
		"11111111F_PBt6CHo3fovdM",
		"111111111_4FzkJ37568tQw",
		"111111111_11jpXCZedGfVR",
		"111111111_111Ahg1opVcGX",
		"111111111_11113CUsUpv9u",
		"111111111_111111VtB5VXd",
		"111111111_11111117YXq9H",
		"111111111_111111112UzHM",
		"111111111_1111111111LUw",
		"111111111_111111111115R",
		"111111111_1111111111112",
	}

	for i := 0; i < 16; i++ {
		uuid := UUID{}
		uuid[i] = byte(1)
		encoded := EncodeBase58(uuid)
		exp := testVecs[i]
		if encoded != exp {
			t.Errorf("%v expected %s, got %s", uuid, exp, encoded)
		}
		dec, err := DecodeBase58(encoded)
		if err != nil {
			t.Error(err)
		}
		if dec != uuid {
			t.Errorf("decoded UUID should be %v, got %v", uuid, dec)
		}
	}
}

func TestOrder(t *testing.T) {
	uuid1 := NewString()
	// sleep for 1 microsecond to make sure the next uuid is different
	time.Sleep(1 * time.Microsecond)
	uuid2 := NewString()

	if uuid1 >= uuid2 {
		t.Errorf("uuid1 should be smaller than uuid2, got %s and %s", uuid1, uuid2)
	}
}

func TestEncodeDecode(t *testing.T) {
	var prev string
	for u0 := 0; u0 < 256; u0++ {
		for u1 := 0; u1 < 256; u1++ {
			for u2 := 0; u2 < 256; u2++ {
				for u3 := 0; u3 < 2; u3++ {

					uuid := UUID{byte(u0), byte(u1), byte(u2), byte(u3), 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, byte(u3 ^ u1)}

					encoded := EncodeBase58(uuid)
					decoded, err := DecodeBase58(encoded)
					if err != nil {
						t.Error(err)
					}
					if decoded != uuid {
						t.Errorf("decoded UUID should be %v, got %v", uuid, decoded)
					}

					if prev != "" && prev >= encoded {
						t.Errorf("prev should be smaller than uuid, got %v and %v", prev, encoded)
					}
					prev = encoded
				}
			}
		}
	}
}

func TestLengthChange22to21(t *testing.T) {
	for lastByte := 0; lastByte < 256; lastByte++ {
		// encodes to 22 characters
		smallerUUID := UUID{0, 0, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, byte(lastByte)}
		smallerEncoded := encodeBase58Raw(smallerUUID)
		if len(smallerEncoded) != 22 {
			t.Errorf("encoded string should be 22 characters long, got %d", len(smallerEncoded))
		}

		// encodes to 21 characters
		largerUUID := UUID{0, 1, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, byte(lastByte)}
		largerEncoded := encodeBase58Raw(largerUUID)
		if len(largerEncoded) != 21 {
			t.Errorf("encoded string should be 21 characters long, got %d", len(largerEncoded))
		}

		// compare without padding - works as expected for all lastByte values
		if smallerEncoded > largerEncoded {
			t.Errorf("smallerEncoded should be smaller than largerEncoded, got %s and %s", smallerEncoded, largerEncoded)
		}

		// compare with padding
		smallerEncodedPad := EncodeBase58(smallerUUID)
		largerEncodedPad := EncodeBase58(largerUUID)
		if smallerEncodedPad > largerEncodedPad {
			t.Errorf("smallerEncodedPad should be smaller than largerEncodedPad, got %s and %s", smallerEncodedPad, largerEncodedPad)
		}
	}
}

func TestLengthChange21to22(t *testing.T) {
	for lastByte := 0; lastByte < 256; lastByte++ {

		// encodes to 21 characters
		smallerUUID := UUID{0, 34, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, byte(lastByte)}
		smallerEncoded := encodeBase58Raw(smallerUUID)
		if len(smallerEncoded) != 21 {
			t.Errorf("encoded string should be 21 characters long, got %d", len(smallerEncoded))
		}

		// encodes to 22 characters
		largerUUID := UUID{0, 35, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, byte(lastByte)}
		largerEncoded := encodeBase58Raw(largerUUID)
		if len(largerEncoded) != 22 {
			t.Errorf("encoded string should be 21 characters long, got %d", len(largerEncoded))
		}

		// compare without padding
		happens := false
		if smallerEncoded > largerEncoded {
			// this happens indeed for all lastByte values
			happens = true
		}

		if !happens {
			t.Errorf("we found counter-example: %v %s and %v %s", smallerUUID, smallerEncoded, largerUUID, largerEncoded)
		}

		// compare with padding
		smallerEncodedPad := EncodeBase58(smallerUUID)
		largerEncodedPad := EncodeBase58(largerUUID)
		if smallerEncodedPad > largerEncodedPad {
			// this never happens for all lastByte values
			t.Errorf("smallerEncodedPad should be smaller than largerEncodedPad, got %s and %s", smallerEncodedPad, largerEncodedPad)
		}

		// last sanity checks
		decSmallerUUID, _ := DecodeBase58(smallerEncodedPad)
		if decSmallerUUID != smallerUUID {
			t.Errorf("decSmallerUUID should be %v, got %v", smallerUUID, decSmallerUUID)
		}

		decLargerUUID, _ := DecodeBase58(largerEncodedPad)
		if decLargerUUID != largerUUID {
			t.Errorf("decLargerUUID should be %v, got %v", largerUUID, decLargerUUID)
		}
	}
}
