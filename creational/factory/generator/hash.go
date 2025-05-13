package generator

import (
	"crypto/sha256"
	"encoding/hex"
)

type HashGenerator struct{}

func (h HashGenerator) Generate(original string) string {
	hash := sha256.Sum256([]byte(original))
	return hex.EncodeToString(hash[:])[:6]
}
