package keeper

import (
	"context"

    "github.com/cdbo/brain/x/membership/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


func (k msgServer) AddMember(goCtx context.Context,  msg *types.MsgAddMember) (*types.MsgAddMemberResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // TODO: Handling the message
    _ = ctx

	return &types.MsgAddMemberResponse{}, nil
}
