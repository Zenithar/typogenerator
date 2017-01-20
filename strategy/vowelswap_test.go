package strategy_test

import (
	"testing"

	"zenithar.org/go/typogenerator/strategy"
)

func TestVowelSwap(t *testing.T) {
	out, err := strategy.VowelSwap.Generate("zenithar.org")
	if err != nil {
		t.Fail()
		t.Fatal("Error should not occurs !", err)
	}

	if len(out) == 0 {
		t.FailNow()
	}

	if len(out) != 18 {
		t.FailNow()
	}
}
