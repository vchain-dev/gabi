package gabi

import (
	"encoding/json"

	"github.com/vchain-dev/gabi/internal/common"
)

// CreateLinkSecret creates json-stringified link secret that we commit to during credential issuance
func CreateLinkSecret(keylength int) string {
	secret, _ := common.RandomBigInt(DefaultSystemParameters[keylength].Lm)
	secretJSON, err := json.Marshal(secret)
	if err != nil {
		return err.Error()
	}

	return string(secretJSON)
}

// CreateNonce creates json-stringified nonce
func CreateNonce(keylength int) string {
	nonce, _ := common.RandomBigInt(DefaultSystemParameters[keylength].Lstatzk)
	nonceJSON, err := json.Marshal(nonce)
	if err != nil {
		return err.Error()
	}

	return string(nonceJSON)
}
