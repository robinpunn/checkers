package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/robinpunn/checkers/testutil/keeper"
	"github.com/robinpunn/checkers/x/checkers/keeper"
	"github.com/robinpunn/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CheckersKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	require.NotNil(t, msgServer)
	require.NotNil(t, context)
}
