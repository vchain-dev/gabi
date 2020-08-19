// +build android ios

package safeprime

import (
	"github.com/vchain-dev/gabi/big"
)

func Generate(int, chan struct{}) (*big.Int, error) {
	panic("Safe prime generation is disabled")
}

func GenerateConcurrent(int, chan struct{}) (<-chan *big.Int, <-chan error) {
	panic("Safe prime generation is disabled")
}
