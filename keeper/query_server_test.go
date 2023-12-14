package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sonrhq/service"
)

func TestQueryParams(t *testing.T) {
	f := initFixture(t)
	require := require.New(t)

	resp, err := f.queryServer.Params(f.ctx, &service.QueryParamsRequest{})
	require.NoError(err)
	require.Equal(service.Params{}, resp.Params)
}

