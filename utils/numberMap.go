package utils

var romanNumberMap = map[string]int32{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func GetMap() map[string]int32 {
	return romanNumberMap
}
