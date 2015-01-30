package modules

import (
	"io"

	"github.com/NebulousLabs/Sia/consensus"
	"github.com/NebulousLabs/Sia/hash"
)

type UploadParams struct {
	Data       io.ReadSeeker
	Duration   consensus.BlockHeight
	Delay      consensus.BlockHeight
	FileSize   uint64
	MerkleRoot hash.Hash

	// these fields are not seen by the host
	Nickname       string
	TotalPieces    int
	RequiredPieces int
	OptimalPieces  int
}

type RentInfo struct {
	Files []string
}

type Renter interface {
	Upload(UploadParams) error
	Download(nickname, filepath string) error
	Rename(currentName, newName string) error
	Info() (RentInfo, error)
}
