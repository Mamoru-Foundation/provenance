package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/provenance-io/provenance/x/name/types"
)

var _ types.QueryServer = Keeper{}

// Params queries params of distribution module
func (k Keeper) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	var params types.Params
	k.paramSpace.GetParamSet(ctx, &params)

	return &types.QueryParamsResponse{Params: params}, nil
}

// Resolve returns the address a name resolves to or an error.
func (k Keeper) Resolve(c context.Context, request *types.QueryResolveRequest) (*types.QueryResolveResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	name, err := k.Normalize(ctx, request.Name)
	if err != nil {
		return nil, err
	}
	record, err := k.GetRecordByName(ctx, name)
	if err != nil {
		return nil, err
	}
	if record == nil {
		return nil, types.ErrNameNotBound
	}
	return &types.QueryResolveResponse{Address: record.Address, Restricted: record.Restricted}, nil
}

// ReverseLookup gets all names bound to an address.
func (k Keeper) ReverseLookup(c context.Context, request *types.QueryReverseLookupRequest) (*types.QueryReverseLookupResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	names := make([]string, 0)
	store := ctx.KVStore(k.storeKey)
	accAddr, err := sdk.AccAddressFromBech32(request.Address)
	if err != nil {
		return nil, types.ErrInvalidAddress
	}
	key, err := types.GetAddressKeyPrefix(accAddr)
	if err != nil {
		return nil, types.ErrInvalidAddress
	}
	nameStore := prefix.NewStore(store, key)
	pageRes, err := query.FilteredPaginate(nameStore, request.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var record types.NameRecord
		err = k.cdc.Unmarshal(value, &record)
		if err != nil {
			return false, err
		}
		if record.Address != request.Address {
			return false, nil
		}
		if accumulate {
			names = append(names, record.Name)
		}
		return true, nil
	})

	if err != nil {
		return nil, err
	}

	return &types.QueryReverseLookupResponse{Name: names, Pagination: pageRes}, nil
}
