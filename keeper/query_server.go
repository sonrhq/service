package keeper

import (
	"context"
	"errors"
	"fmt"

	"cosmossdk.io/collections"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sonrhq/service"
)

var _ service.QueryServer = queryServer{}

// NewQueryServerImpl returns an implementation of the module QueryServer.
func NewQueryServerImpl(k Keeper) service.QueryServer {
	return queryServer{k}
}

type queryServer struct {
	k Keeper
}

// Counter defines the handler for the Query/Counter RPC method.
func (qs queryServer) Counter(ctx context.Context, req *service.QueryCounterRequest) (*service.QueryCounterResponse, error) {
	if _, err := qs.k.addressCodec.StringToBytes(req.Address); err != nil {
		return nil, fmt.Errorf("invalid sender address: %w", err)
	}

	counter, err := qs.k.Counter.Get(ctx, req.Address)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return &service.QueryCounterResponse{Counter: 0}, nil
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &service.QueryCounterResponse{Counter: counter}, nil
}

// Counters defines the handler for the Query/Counters RPC method.
func (qs queryServer) Counters(ctx context.Context, req *service.QueryCountersRequest) (*service.QueryCountersResponse, error) {
	counters, pageRes, err := query.CollectionPaginate(
		ctx,
		qs.k.Counter,
		req.Pagination,
		func(key string, value uint64) (*service.Counter, error) {
			return &service.Counter{
				Address: key,
				Count:   value,
			}, nil
		})
	if err != nil {
		return nil, err
	}

	return &service.QueryCountersResponse{Counters: counters, Pagination: pageRes}, nil
}

// Params defines the handler for the Query/Params RPC method.
func (qs queryServer) Params(ctx context.Context, req *service.QueryParamsRequest) (*service.QueryParamsResponse, error) {
	params, err := qs.k.Params.Get(ctx)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return &service.QueryParamsResponse{Params: service.Params{}}, nil
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &service.QueryParamsResponse{Params: params}, nil
}
