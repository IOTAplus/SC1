package test

import (
	"testing"

	"github.com/iotaledger/wasp/tools/schema/crowdsale"
	"github.com/iotaledger/wasp/packages/vm/wasmsolo"
	"github.com/stretchr/testify/require"
)

func TestDeploy(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, crowdsale.ScName, crowdsale.OnLoad)
	require.NoError(t, ctx.ContractExists(crowdsale.ScName))
}
