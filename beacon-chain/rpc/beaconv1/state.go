package beaconv1

import (
	"bytes"
	"context"
	"errors"

	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1"
	"github.com/prysmaticlabs/prysm/shared/params"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// GetGenesis retrieves details of the chain's genesis which can be used to identify chain.
func (bs *Server) GetGenesis(ctx context.Context, _ *emptypb.Empty) (*ethpb.GenesisResponse, error) {
	ctx, span := trace.StartSpan(ctx, "beaconv1.GetGenesis")
	defer span.End()

	genesisTime := bs.GenesisTimeFetcher.GenesisTime()
	if genesisTime.IsZero() {
		return nil, status.Errorf(codes.NotFound, "Chain genesis info is not yet known")
	}
	validatorRoot := bs.ChainInfoFetcher.GenesisValidatorRoot()
	if bytes.Equal(validatorRoot[:], params.BeaconConfig().ZeroHash[:]) {
		return nil, status.Errorf(codes.NotFound, "Chain genesis info is not yet known")
	}
	forkVersion := params.BeaconConfig().GenesisForkVersion

	return &ethpb.GenesisResponse{
		Data: &ethpb.Genesis{
			GenesisTime: &timestamppb.Timestamp{
				Seconds: genesisTime.Unix(),
				Nanos:   0,
			},
			GenesisValidatorsRoot: validatorRoot[:],
			GenesisForkVersion:    forkVersion,
		},
	}, nil
}

// GetStateRoot calculates HashTreeRoot for state with given 'stateId'. If stateId is root, same value will be returned.
func (bs *Server) GetStateRoot(ctx context.Context, req *ethpb.StateRequest) (*ethpb.StateRootResponse, error) {
	return nil, errors.New("unimplemented")
}

// GetStateFork returns Fork object for state with given 'stateId'.
func (bs *Server) GetStateFork(ctx context.Context, req *ethpb.StateRequest) (*ethpb.StateForkResponse, error) {
	return nil, errors.New("unimplemented")
}

// GetFinalityCheckpoints returns finality checkpoints for state with given 'stateId'. In case finality is
// not yet achieved, checkpoint should return epoch 0 and ZERO_HASH as root.
func (bs *Server) GetFinalityCheckpoints(ctx context.Context, req *ethpb.StateRequest) (*ethpb.StateFinalityCheckpointResponse, error) {
	return nil, errors.New("unimplemented")
}
