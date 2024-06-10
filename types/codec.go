package types

import "github.com/cosmos/cosmos-sdk/codec"

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSignData{}, "sign/MsgSignData", nil)
}

func init() {
	RegisterLegacyAminoCodec(amino)
	amino.Seal()
}
