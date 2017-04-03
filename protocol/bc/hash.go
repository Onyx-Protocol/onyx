package bc

import (
	"database/sql/driver"
	"encoding/binary"
	"io"

	"golang.org/x/crypto/sha3"

	"chain/crypto/sha3pool"
	"chain/encoding/blockchain"
)

var EmptyStringHash Hash

func init() {
	EmptyStringHash.FromByte32(sha3.Sum256(nil))
}

// Hash represents a 256-bit hash.  Data structure defined in
// hash.proto.

func (h Hash) Byte32() (b32 Byte32) {
	binary.BigEndian.PutUint64(b32[0:8], h.V0)
	binary.BigEndian.PutUint64(b32[8:16], h.V1)
	binary.BigEndian.PutUint64(b32[16:24], h.V2)
	binary.BigEndian.PutUint64(b32[24:32], h.V3)
	return b32
}

func (h *Hash) FromByte32(b32 Byte32) {
	h.V0 = binary.BigEndian.Uint64(b32[0:8])
	h.V1 = binary.BigEndian.Uint64(b32[8:16])
	h.V2 = binary.BigEndian.Uint64(b32[16:24])
	h.V3 = binary.BigEndian.Uint64(b32[24:32])
}

func (h *Hash) FromHasher(hasher sha3.ShakeHash) {
	var b32 Byte32
	hasher.Read(b32[:])
	h.FromByte32(b32)
}

// MarshalText satisfies the TextMarshaler interface.
// It returns the bytes of h encoded in hex,
// for formats that can't hold arbitrary binary data.
// It never returns an error.
func (h Hash) MarshalText() ([]byte, error) {
	return h.Byte32().MarshalText()
}

// UnmarshalText satisfies the TextUnmarshaler interface.
// It decodes hex data from b into h.
func (h *Hash) UnmarshalText(b []byte) error {
	var b32 Byte32
	err := b32.UnmarshalText(b)
	if err != nil {
		return err
	}
	h.FromByte32(b32)
	return nil
}

// UnmarshalJSON satisfies the json.Unmarshaler interface.
// If b is a JSON-encoded null, it copies the zero-value into h. Othwerwise, it
// decodes hex data from b into h.
func (h *Hash) UnmarshalJSON(b []byte) error {
	var b32 Byte32
	err := b32.UnmarshalJSON(b)
	if err != nil {
		return err
	}
	h.FromByte32(b32)
	return nil
}

func (h Hash) Bytes() []byte {
	b32 := h.Byte32()
	return b32[:]
}

// Value satisfies the driver.Valuer interface
func (h Hash) Value() (driver.Value, error) {
	return h.Bytes(), nil
}

// Scan satisfies the driver.Scanner interface
func (h *Hash) Scan(val interface{}) error {
	var b32 Byte32
	err := b32.Scan(val)
	if err != nil {
		return err
	}
	h.FromByte32(b32)
	return nil
}

func writeFastHash(w io.Writer, d []byte) error {
	if len(d) == 0 {
		_, err := blockchain.WriteVarstr31(w, nil)
		return err
	}
	var h [32]byte
	sha3pool.Sum256(h[:], d)
	_, err := blockchain.WriteVarstr31(w, h[:])
	return err
}

// WriteTo writes p to w.
func (h *Hash) WriteTo(w io.Writer) (int64, error) {
	return h.Byte32().WriteTo(w)
}

func (h *Hash) readFrom(r io.Reader) (int, error) {
	var b32 Byte32
	n, err := b32.ReadFrom(r)
	if err != nil {
		return int(n), err
	}
	h.FromByte32(b32)
	return int(n), nil
}
