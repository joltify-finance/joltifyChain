package keeper

import (
	"context"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/joltify/joltifychain/joltifychain/x/vault/types"
)

func (k msgServer) CreateCreatePool(goCtx context.Context, msg *types.MsgCreateCreatePool) (*types.MsgCreateCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	height, err := strconv.ParseInt(msg.BlockHeight, 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("invalid block height %v", msg.BlockHeight))
	}
	history, get := k.vaultStaking.GetHistoricalInfo(ctx, height)
	if !get {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("too early, we cannot find the block %v", msg.BlockHeight))
	}

	// now we check whether the msg is sent from the validator
	validators := history.GetValset()
	isValidator := false
	for _, el := range validators {
		if el.GetOperator().Equals(msg.Creator) {
			isValidator = true
			break
		}
	}
	if !isValidator {
		ctx.Logger().Info("not a validator update tss message", "result", "false")
		return &types.MsgCreateCreatePoolResponse{Successful: false}, nil
	}
	var newProposal types.PoolProposal
	info, isFound := k.GetCreatePool(ctx, msg.BlockHeight)
	if isFound {
		entryFound := false
		for i, proposal := range info.Proposal {
			newProposal.PoolPubKey = proposal.PoolPubKey
			if proposal.GetPoolPubKey() == msg.PoolPubKey {
				proposal.Nodes = append(proposal.Nodes, msg.Creator)
				entryFound = true
				info.Proposal[i] = proposal
				break
			}
		}
		if !entryFound {
			newProposal.PoolPubKey = msg.PoolPubKey
			newProposal.Nodes = []sdk.AccAddress{msg.Creator}
			info.Proposal = append(info.Proposal, &newProposal)
		}

	} else {
		pro := types.PoolProposal{
			PoolPubKey: msg.PoolPubKey,
			Nodes:      []sdk.AccAddress{msg.Creator},
		}
		info.Proposal = []*types.PoolProposal{&pro}
	}
	var createPool = types.CreatePool{
		BlockHeight: msg.BlockHeight,
		Validators:  validators,
		Proposal:    info.Proposal,
	}

	k.SetCreatePool(
		ctx,
		createPool,
	)

	return &types.MsgCreateCreatePoolResponse{Successful: true}, nil
}
