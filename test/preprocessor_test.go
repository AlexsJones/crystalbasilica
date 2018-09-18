package test

import (
	"github.com/AlexsJones/crystalbasilica/src/loader"
	"github.com/AlexsJones/crystalbasilica/src/preprocessor"
	"testing"
)

func TestSHAGeneration(t *testing.T) {

	fl := &loader.FileLoader{}

	basilica, err := loader.ILoader.LoadFile(fl,"files/test-01.yaml")
	if err != nil {
		t.Error(err)
	}

	testSHA := []string{ "dd8f114f75bc022718f2e03136f55e5ca8dff0a8","bc1260facaf3eb3b977185ba4f50d468e53568a0" }

	sp := &preprocessor.StandardPreprocessor{}

	for c,i := range basilica.Operations {

		sha, _ := preprocessor.IPreprocessor.CalculateSHA(sp, i.Directive.Command)
		if testSHA[c] != sha {
			t.Error("Bad SHA")
		}
	}

}