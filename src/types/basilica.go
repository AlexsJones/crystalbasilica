package types

type Basilica struct {
	Version string `yaml:"Version"`
	Operations []struct {
		Directive struct {
			Command string `yaml:"Command"`
			Expectation string `yaml:"Expectation"`
		}`yaml:"Directive"`
	}`yaml:"Operations"`
}