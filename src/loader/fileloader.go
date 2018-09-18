package loader

import (
	"github.com/AlexsJones/crystalbasilica/src/types"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type FileLoader struct {

}

func (f *FileLoader) LoadFile(path string) (*types.Basilica, error) {

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	c := types.Basilica{}
	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		log.Printf("Failed to load path: %s\n", path)
		return nil, err
	}
	return &c, nil
}