package preprocessor


type IPreprocessor interface {
	CalculateSHA(input string) (string,error)
}

func CalculateSHA(ip IPreprocessor, input string) (string,error) {
	return ip.CalculateSHA(input)
}