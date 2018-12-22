package crypto_test

import (
	"testing"
)

func Test_Hash(t *testing.T) {
	t.Run("Can hash and compare", shouldBeAbleToHashAndCompareStrings)
	t.Run("Can detect unequal hashes", shouldBeErrorWhenComparingUnequalHash)
	t.Run("Generate a different salt every run", shouldGenerateDifferentSaltEveryRun)
}

func shouldBeAbleToHashAndCompareStrings(t *testing.T) {

}
