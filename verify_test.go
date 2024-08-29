// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package adr36_test

import (
	"encoding/base64"
	"testing"

	"adr36.dev"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestVerifySignature(t *testing.T) {
	// The below signature was generated using Keplr's signArbitrary function.
	//
	// share bubble good swarm sustain leaf burst build spirit inflict undo shadow antique warm soft praise foam slab laptop hint giggle also book treat
	//
	// {
	//     "pub_key": {
	//         "type": "tendermint/PubKeySecp256k1",
	//         "value": "AlE8CxHR19ID5lxrVtTxSgJFlK3T+eYtyDM/vBA3Fowr"
	//     },
	//     "signature": "qe5dDxdOgY8B2LjMqnK5/5iRIFOCwdTu0G5ZQ66bHzVgP15V2Fb+fzOH0wPAUC5GUQ23M1cSvysulzKIbXY/4Q=="
	// }

	bz, _ := base64.StdEncoding.DecodeString("AlE8CxHR19ID5lxrVtTxSgJFlK3T+eYtyDM/vBA3Fowr")
	signature, _ := base64.StdEncoding.DecodeString("qe5dDxdOgY8B2LjMqnK5/5iRIFOCwdTu0G5ZQ66bHzVgP15V2Fb+fzOH0wPAUC5GUQ23M1cSvysulzKIbXY/4Q==")

	valid := adr36.VerifySignature(&secp256k1.PubKey{Key: bz}, []byte("I hereby declare that I am the address owner."), signature)
	require.True(t, valid)
}

func init() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("noble", "noblepub")
	config.Seal()
}
