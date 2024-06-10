package adr36

import (
	"adr36.dev/types"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
)

func VerifySignature(pubkey cryptotypes.PubKey, data []byte, signature []byte) bool {
	bz := legacytx.StdSignBytes(
		"",
		0,
		0,
		0,
		legacytx.NewStdFee(0, sdk.Coins{}), //nolint:staticcheck
		[]sdk.Msg{&types.MsgSignData{
			Signer: sdk.AccAddress(pubkey.Address()).String(),
			Data:   data,
		}},
		"",
	)

	return pubkey.VerifySignature(bz, signature)
}
