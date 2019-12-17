package codec

// URLCodec - URL codec structure with bound methods
type URLCodec struct {
	startValue int64
	codec      GenericCodec
}

var (
	// StartValue - initial offset for URL codec to make shortened URLs nice
	StartValue = int64(3844) // nolint
	// Dictionary - URL codec dictionary, b62 at the moment
	Dictionary = []string{ // nolint
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
		"k", "l", "m", "n", "o", "p", "q", "r", "s", "t",
		"u", "v", "w", "x", "y", "z", "A", "B", "C", "D",
		"E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X",
		"Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7",
		"8", "9",
	}
)

// Encode - translate URL to its unique integer id
func (u *URLCodec) Encode(requestURI string) int64 {
	return u.codec.DicToDec(requestURI) - u.startValue
}

// Decode - translate unique URL id to its shortened version
func (u *URLCodec) Decode(encoded int64) string {
	return u.codec.DecToDic(encoded + u.startValue)
}

// NewURLCodec - helper function that creates new URLCodec instance with supplied initial offset and dictionary
func NewURLCodec(startValue int64, dictionary []string) *URLCodec {
	return &URLCodec{
		startValue: startValue,
		codec:      NewCodec(dictionary),
	}
}

// New - helper function that creates and populates URLCodec instance with default values
func New() *URLCodec {
	return NewURLCodec(StartValue, Dictionary)
}
