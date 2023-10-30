package uuid7

import (
	"fmt"
	"testing"
	"time"
)

func toBinary(u UUID) string {
	var binaryStr string
	for _, b := range u {
		binaryStr += fmt.Sprintf("%08b", b)
	}
	return binaryStr
}

func TestUUID7(t *testing.T) {
	prev := ""
	for i := 0; i < 120; i++ {
		uTime := time.Now()

		uuid, err := FromTime(uTime)
		if err != nil {
			t.Error(err)
		}
		binStr := toBinary(uuid)
		/*
		   unix_ts_ms:
		      48 bit big-endian unsigned number of Unix epoch timestamp in
		      milliseconds as per Section 6.1.  Occupies bits 0 through 47
		      (octets 0-5).

		   ver:
		      The 4 bit version field as defined by Section 4.2, set to 0b0111
		      (7).  Occupies bits 48 through 51 of octet 6.

		   rand_a:
		      12 bits pseudo-random data to provide uniqueness as per
		      Section 6.8 and/or optional constructs to guarantee additional
		      monotonicity as per Section 6.2.  Occupies bits 52 through 63
		      (octets 6-7).

		   var:
		      The 2 bit variant field as defined by Section 4.1, set to 0b10.
		      Occupies bits 64 and 65 of octet 8.

		   rand_b:
		      The final 62 bits of pseudo-random data to provide uniqueness as
		      per Section 6.8 and/or an optional counter to guarantee additional
		      monotonicity as per Section 6.2.  Occupies bits 66 through 127
		      (octets 8-15).
		*/

		// compare unix_ts_ms
		unixTsMs := fmt.Sprintf("%048b", uint64(uTime.UnixNano()/int64(time.Millisecond)))
		if binStr[0:48] != unixTsMs {
			t.Errorf("unix_ts_ms bits should be %s, got %s", unixTsMs, binStr[0:48])
		}

		// check ver
		if binStr[48:52] != "0111" {
			t.Errorf("ver bits should be 0111, got %s", binStr[48:52])
		}

		// check var
		if binStr[64:66] != "10" {
			t.Errorf("var bits should be 10, got %s", binStr[64:66])
		}

		time.Sleep(1 * time.Millisecond)

		if i > 0 && binStr <= prev {
			t.Errorf("uuid should be greater than previous, got %s and %s", binStr, prev)
		}
		prev = binStr
	}

}
