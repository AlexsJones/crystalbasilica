package loader

import "github.com/AlexsJones/crystalbasilica/src/types"

type ILoader interface {
	LoadFile(path string) (*types.Basilica, error)
}

func LoadFile(il ILoader, path string) (*types.Basilica, error) {
	return il.LoadFile(path)
}