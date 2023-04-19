package hashstr

import (
	"fmt"
	"testing"
	"crypto/sha256"
)

func TestSha256ToString(t *testing.T) {
	testStr := "This is a test"
	if fmt.Sprintf("%x", sha256.Sum256([]byte(testStr))) != Sha256ToString(testStr) {
		t.Errorf("Invalid result for testStr = %s\n", testStr)
	}
}
