package test

import (
	"github.com/AlexsJones/crystalbasilica/src/loader"
	"testing"
)

func TestFileLoad(t *testing.T) {

	fl := &loader.FileLoader{}

	_, err := loader.ILoader.LoadFile(fl,"files/test-01.yaml")
	if err != nil {
		t.Error(err)
	}

}

func TestMissingFileLoad (t *testing.T) {
	fl := &loader.FileLoader{}

	_, err := loader.ILoader.LoadFile(fl,"files/test-missing.yaml")
	if err == nil {
		t.Error(err)
	}
}