// Package iface defines an interface for the validator database.
package iface

import (
	"context"
	"io"

	"github.com/prysmaticlabs/prysm/validator/db/kv"
)

// ValidatorDB defines the necessary methods for a Prysm validator DB.
type ValidatorDB interface {
	io.Closer
	DatabasePath() string
	ClearDB() error
	UpdatePublicKeysBuckets(publicKeys [][48]byte) error

	// Genesis information related methods.
	GenesisValidatorsRoot(ctx context.Context) ([]byte, error)
	SaveGenesisValidatorsRoot(ctx context.Context, genValRoot []byte) error

	// New data structure methods
	ProposalHistoryForSlot(ctx context.Context, publicKey []byte, slot uint64) ([]byte, error)
	SaveProposalHistoryForSlot(ctx context.Context, pubKey []byte, slot uint64, signingRoot []byte) error
	SaveProposalHistoryForPubKeysV2(ctx context.Context, proposals map[[48]byte]kv.ProposalHistoryForPubkey) error

	// New attestation store methods
	AttestationHistoryForPubKeysV2(ctx context.Context, publicKeys [][48]byte) (map[[48]byte]kv.EncHistoryData, error)
	AttestationHistoryForPubKeyV2(ctx context.Context, publicKey [48]byte) (kv.EncHistoryData, error)
	SaveAttestationHistoryForPubKeysV2(ctx context.Context, historyByPubKeys map[[48]byte]kv.EncHistoryData) error
	SaveAttestationHistoryForPubKeyV2(ctx context.Context, pubKey [48]byte, history kv.EncHistoryData) error
}
