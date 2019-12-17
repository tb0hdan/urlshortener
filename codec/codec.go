package codec

import (
	"math"
	"strings"
)

// Codec - data codec definition
type Codec struct {
	dictionary []string
	dicLenf64  float64
	dicLen64   int64
}

func (c *Codec) index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}

	return -1
}

// DecToDic - convert decimal value to string
func (c *Codec) DecToDic(dec int64) string {
	resultSlice := make([]string, 0)

	for dec > 0 {
		quotient, remainder := dec/c.dicLen64, dec%c.dicLen64

		resultSlice = append(resultSlice, "")
		copy(resultSlice[1:], resultSlice[0:])
		resultSlice[0] = c.dictionary[remainder]
		dec = quotient
	}

	return strings.Join(resultSlice, "")
}

// DicToDec - convert string to decimal value
func (c *Codec) DicToDec(dic string) int64 {
	var (
		b10 int64
	)

	x := len(dic) - 1

	for _, ch := range strings.Split(dic, "") {
		val := int64(c.index(c.dictionary, ch))
		b10 += val * int64(math.Pow(c.dicLenf64, float64(x)))
		x--
	}

	return b10
}

// NewCodec - helper function that creates new codec instance with supplied dictionary
func NewCodec(dictionary []string) *Codec {
	return &Codec{
		dictionary: dictionary,
		dicLenf64:  float64(len(dictionary)),
		dicLen64:   int64(len(dictionary)),
	}
}
