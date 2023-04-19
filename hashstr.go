package hashstr

import (
	"fmt"
	"crypto/sha256"
)

func Sha256ToString(contents interface{}) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(contents.(string))))
}
