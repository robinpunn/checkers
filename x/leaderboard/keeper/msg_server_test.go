package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/robinpunn/checkers/testutil/keeper"
	"github.com/robinpunn/checkers/x/leaderboard/keeper"
	"github.com/robinpunn/checkers/x/leaderboard/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.LeaderboardKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
