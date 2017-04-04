package bc

import (
	"database/sql/driver"
	"io"

	"chain/crypto/sha3pool"
	"chain/encoding/blockchain"
)

// AssetID is the Hash256 of the issuance script for the asset and the
// initial block of the chain where it appears.
type AssetID [32]byte

func (a AssetID) Bytes() []byte { return a[:] }

func (a AssetID) String() string                { return Hash(a).String() }
func (a AssetID) MarshalText() ([]byte, error)  { return Hash(a).MarshalText() }
func (a *AssetID) UnmarshalText(b []byte) error { return (*Hash)(a).UnmarshalText(b) }
func (a *AssetID) UnmarshalJSON(b []byte) error { return (*Hash)(a).UnmarshalJSON(b) }
func (a AssetID) Value() (driver.Value, error)  { return Hash(a).Value() }
func (a *AssetID) Scan(b interface{}) error     { return (*Hash)(a).Scan(b) }

func (ad AssetDefinition) ComputeAssetID() (assetID AssetID) {
	h := sha3pool.Get256()
	defer sha3pool.Put256(h)
	writeForHash(h, ad) // error is impossible
	h.Read(assetID[:])
	return assetID
}

func ComputeAssetID(prog []byte, initialBlockID Hash, vmVersion uint64, data Hash) AssetID {
	def := &AssetDefinition{
		InitialBlockId: initialBlockID.Proto(),
		IssuanceProgram: &Program{
			VmVersion: vmVersion,
			Code:      prog,
		},
		Data: data.Proto(),
	}
	return def.ComputeAssetID()
}

type AssetAmount struct {
	AssetID AssetID `json:"asset_id"`
	Amount  uint64  `json:"amount"`
}

func (a *AssetAmount) readFrom(r io.Reader) (int, error) {
	n1, err := io.ReadFull(r, a.AssetID[:])
	if err != nil {
		return n1, err
	}
	var n2 int
	a.Amount, n2, err = blockchain.ReadVarint63(r)
	return n1 + n2, err
}

func (a *AssetAmount) writeTo(w io.Writer) error {
	_, err := w.Write(a.AssetID[:])
	if err != nil {
		return err
	}
	_, err = blockchain.WriteVarint63(w, a.Amount)
	return err
}

func (pa *ProtoAssetID) AssetID() AssetID {
	if pa == nil {
		return AssetID{}
	}
	h := pa.Hash.Hash()
	return AssetID(h)
}

func (a AssetID) Proto() *ProtoAssetID {
	return &ProtoAssetID{
		Hash: Hash(a).Proto(),
	}
}

func (paa *ProtoAssetAmount) AssetAmount() AssetAmount {
	if paa == nil {
		return AssetAmount{}
	}

	return AssetAmount{
		AssetID: paa.AssetId.AssetID(),
		Amount:  paa.Amount,
	}
}

func (aa AssetAmount) Proto() *ProtoAssetAmount {
	return &ProtoAssetAmount{
		AssetId: aa.AssetID.Proto(),
		Amount:  aa.Amount,
	}
}
