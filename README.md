# Hypersequent's UUID7

This is a Go library for generating UUIDv7 based 
on [draft 11 of RFC4122bis](https://datatracker.ietf.org/doc/draft-ietf-uuidrev-rfc4122bis/11/) — draft of the new UUID specification.

Besides this library implements a custom string encoding for UUIDv7, which is lexicographically sortable. This string 
encoding is not defined in the RFC4122bis draft. The string encoding is based on Base58 encoding used in Bitcoin.

## Binary format

```plain 
    0                   1                   2                   3
    0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |                           unix_ts_ms                          |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |          unix_ts_ms           |  ver  |       rand_a          |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |var|                        rand_b                             |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |                            rand_b                             |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
```

- **unix_ts_ms** is filled with Go's `time.Now().UnixNano() / 1e6`
- **ver** is `0b0111` for UUIDv7 (RFC4122bis draft 11)
- **rand_a** is filled using "Replace Left-Most Random Bits with Increased Clock Precision" (Method 3)
- **var** is `0b10` for UUIDv7 (RFC4122bis draft 11)
- **rand_b** is cryptographically random bits, generated using Go's `crypto/rand` package

## String Encoding 

The UUID is encoded using Base58 encoding using BTC alphabet, which is the same as the one used in [Bitcoin](https://en.bitcoinwiki.org/wiki/Base58). 
Bitcoin address checksum is not used. 

The encoded string is always 22 characters long (padded with leading "zero" digit `1` if needed).

String representation is sortable lexicographically, which is good for databases.


## Usage

```go
ustr := hyperuuid7.NewString() // returns a 22 character long string like "1C3Rttz29K2U2o4AdhPF5b"
```
## Dependencies

- [github.com/mr-tron/base58](https://github.com/mr-tron/base58) - Base58 encoding/decoding package (MIT License)

## License
MIT 


