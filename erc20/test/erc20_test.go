package test

import (
	"testing"

	"github.com/iotaplus/SC1/erc20/go/erc20"
	"github.com/iotaledger/wasp/packages/vm/wasmsolo"
	"github.com/stretchr/testify/require"
)

func TestDeploy(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, erc20.ScName, erc20.OnLoad)
	require.NoError(t, ctx.ContractExists(erc20.ScName))
}
