package beaconv1

import (
	"context"
	"testing"
	"time"

	chainMock "github.com/prysmaticlabs/prysm/beacon-chain/blockchain/testing"
	"github.com/prysmaticlabs/prysm/shared/params"
	"github.com/prysmaticlabs/prysm/shared/testutil/assert"
	"github.com/prysmaticlabs/prysm/shared/testutil/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestGetGenesis(t *testing.T) {
	params.SetupTestConfigCleanup(t)
	config := params.BeaconConfig()
	config.GenesisForkVersion = []byte("genesis")
	params.OverrideBeaconConfig(config)
	genesis := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	validatorsRoot := [32]byte{1, 2, 3, 4, 5, 6}

	t.Run("OK", func(t *testing.T) {
		chainService := &chainMock.ChainService{
			Genesis:        genesis,
			ValidatorsRoot: validatorsRoot,
		}
		s := Server{
			GenesisTimeFetcher: chainService,
			ChainInfoFetcher:   chainService,
		}
		resp, err := s.GetGenesis(context.Background(), &emptypb.Empty{})
		require.NoError(t, err)
		assert.Equal(t, genesis.Unix(), resp.Data.GenesisTime.Seconds)
		assert.Equal(t, int32(0), resp.Data.GenesisTime.Nanos)
		assert.DeepEqual(t, validatorsRoot[:], resp.Data.GenesisValidatorsRoot)
		assert.DeepEqual(t, []byte("genesis"), resp.Data.GenesisForkVersion)
	})

	t.Run("No genesis time", func(t *testing.T) {
		chainService := &chainMock.ChainService{
			Genesis:        time.Time{},
			ValidatorsRoot: validatorsRoot,
		}
		s := Server{
			GenesisTimeFetcher: chainService,
			ChainInfoFetcher:   chainService,
		}
		_, err := s.GetGenesis(context.Background(), &emptypb.Empty{})
		assert.ErrorContains(t, "Chain genesis info is not yet known", err)
	})

	t.Run("No genesis validator root", func(t *testing.T) {
		chainService := &chainMock.ChainService{
			Genesis:        genesis,
			ValidatorsRoot: [32]byte{},
		}
		s := Server{
			GenesisTimeFetcher: chainService,
			ChainInfoFetcher:   chainService,
		}
		_, err := s.GetGenesis(context.Background(), &emptypb.Empty{})
		assert.ErrorContains(t, "Chain genesis info is not yet known", err)
	})
}
