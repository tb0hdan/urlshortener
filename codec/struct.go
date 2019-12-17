package codec

// GenericCodec - Definition for a simple back-and-forth codec
// Dictionary can be supplied to either New(dictionary []string) (if any) or directly to the structure
type GenericCodec interface {
	// DecToDic - convert decimal value to string
	DecToDic(dec int64) string
	// DicToDec - convert string to decimal value
	DicToDec(dic string) int64
}
