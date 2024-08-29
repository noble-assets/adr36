// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

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
