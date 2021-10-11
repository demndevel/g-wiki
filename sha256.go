package main

import (
	"crypto/sha256"
	"encoding/hex"
)

func encode(toencode string) string {
	h := sha256.New()
	h.Write([]byte(toencode))
	return hex.EncodeToString(h.Sum(nil))
}
