package preprocessor

import (
	"crypto/sha1"
	"encoding/hex"
)

type StandardPreprocessor struct {

}
func (sp *StandardPreprocessor) CalculateSHA( input string) (string,error) {
	h := sha1.New()
	h.Write([]byte(input))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash, nil
}