package wormhole_ibc

import (
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/stratisproject/wormhole/sdk/vaa"
	"github.com/stretchr/testify/require"
	"github.com/wormhole-foundation/wormchain/interchaintest/guardians"
	"github.com/wormhole-foundation/wormchain/interchaintest/helpers"
)

func SubmitWormholeIbcUpdateChannelChainMsg(t *testing.T,
	allowlistChainID vaa.ChainID,
	allowlistChannel string,
	guardians *guardians.ValSet) string {

	paddedChannel, _ := vaa.LeftPadIbcChannelId(allowlistChannel)

	bodyIbcReceiverUpdateChannelChain := vaa.BodyIbcUpdateChannelChain{
		TargetChainId: vaa.ChainIDWormchain,
		ChannelId:     paddedChannel,
		ChainId:       allowlistChainID,
	}

	payload, err := bodyIbcReceiverUpdateChannelChain.Serialize(vaa.IbcReceiverModuleStr)
	require.NoError(t, err)

	v := helpers.GenerateGovernanceVaa(0, guardians, payload)
	vBz, err := v.Marshal()
	require.NoError(t, err)
	vHex := base64.StdEncoding.EncodeToString(vBz)

	submitVAAMsg := ExecuteMsg{
		SubmitVAA:                nil,
		PostMessage:              nil,
		SubmitUpdateChannelChain: &ExecuteMsg_SubmitUpdateChannelChain{Vaa: Binary(vHex)},
	}

	submitVAAMsgBz, err := json.Marshal(submitVAAMsg)
	require.NoError(t, err)

	return string(submitVAAMsgBz)
}
