package test

import (
	"testing"

	"github.com/iotaplus/SC1/crowdsale/go/ico"
	"github.com/iotaledger/wasp/packages/vm/wasmsolo"
	"github.com/stretchr/testify/require"
)

func TestDeploy(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, ico.ScName, ico.OnLoad)
	require.NoError(t, ctx.ContractExists(ico.ScName))
}
