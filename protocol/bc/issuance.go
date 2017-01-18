package bc

import "chain/crypto/sha3pool"

type IssuanceInput struct {
	// Commitment
	Nonce  []byte
	Amount uint64
	// Note: as long as we require serflags=0x7, we don't need to
	// explicitly store the asset ID here even though it's technically
	// part of the input commitment. We can compute it instead from
	// values in the witness (which, with serflags other than 0x7,
	// might not be present).

	// Witness
	InitialBlock    Hash
	AssetDefinition []byte
	VMVersion       uint64
	IssuanceProgram []byte
	Arguments       [][]byte
}

func (ii *IssuanceInput) IsIssuance() bool { return true }

func (ii *IssuanceInput) AssetID() AssetID {
	return ComputeAssetID(ii.IssuanceProgram, ii.InitialBlock, ii.VMVersion, ii.AssetDefinitionHash())
}

func (ii *IssuanceInput) AssetDefinitionHash() (defhash Hash) {
	sha := sha3pool.Get256()
	defer sha3pool.Put256(sha)
	sha.Write(ii.AssetDefinition)
	sha.Read(defhash[:])
	return
}

func NewIssuanceInput(
	nonce []byte,
	amount uint64,
	referenceData []byte,
	initialBlock Hash,
	issuanceProgram []byte,
	arguments [][]byte,
	assetDefinition []byte,
) *TxInput {
	return &TxInput{
		AssetVersion:  1,
		ReferenceData: referenceData,
		TypedInput: &IssuanceInput{
			Nonce:           nonce,
			Amount:          amount,
			InitialBlock:    initialBlock,
			AssetDefinition: assetDefinition,
			VMVersion:       1,
			IssuanceProgram: issuanceProgram,
			Arguments:       arguments,
		},
	}
}
