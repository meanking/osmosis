package keeper

import (
	"fmt"

	"github.com/osmosis-labs/osmosis/v11/x/incentives/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)
	k.SetLockableDurations(ctx, genState.LockableDurations)
	for _, gauge := range genState.Gauges {
		if gauge.Id == 1894 {
			fmt.Printf("ADAM %v \n", gauge)
		}
		err := k.SetGaugeWithRefKey(ctx, &gauge)
		if err != nil {
			panic(err)
		}
	}
	k.SetLastGaugeID(ctx, genState.LastGaugeId)
}

// ExportGenesis returns the capability module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return &types.GenesisState{
		Params:            k.GetParams(ctx),
		LockableDurations: k.GetLockableDurations(ctx),
		Gauges:            k.GetNotFinishedGauges(ctx),
		LastGaugeId:       k.GetLastGaugeID(ctx),
	}
}
