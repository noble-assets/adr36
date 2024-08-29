// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
)

var _ legacytx.LegacyMsg = &MsgSignData{}

func (msg *MsgSignData) ValidateBasic() error {
	// TODO: Implement this ...
	return nil
}

func (msg *MsgSignData) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{signer}
}

func (msg *MsgSignData) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (*MsgSignData) Route() string { return "sign" }

func (*MsgSignData) Type() string { return "MsgSignData" }
