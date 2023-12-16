package keeper

import (
	"github.com/robinpunn/checkers/x/leaderboard/types"
)

var _ types.QueryServer = Keeper{}
