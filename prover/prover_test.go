package prover

import (
	"context"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	log.Root().SetHandler(
		log.LvlFilterHandler(
			log.LvlDebug,
			log.StreamHandler(os.Stdout, log.TerminalFormat(true)),
		),
	)
	os.Exit(m.Run())
}

func newTestProver(t *testing.T) *Prover {
	l1ProverPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROVER_PRIVATE_KEY")))
	require.Nil(t, err)

	p := new(Prover)

	require.Nil(t, initFromConfig(p, &Config{
		L1Endpoint:      os.Getenv("L1_NODE_ENDPOINT"),
		L2Endpoint:      os.Getenv("L2_NODE_ENDPOINT"),
		TaikoL1Address:  common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:  common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		L1ProverPrivKey: l1ProverPrivKey,
		Dummy:           true,
	}))

	return p
}

func TestOnForceTimer(t *testing.T) {
	p := newTestProver(t)

	err := p.onForceTimer(context.Background())

	require.Nil(t, err)
}