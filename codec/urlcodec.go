package codec

type URLCodec struct {
	startValue int64
	codec      *Codec
}

var (
	StartValue = int64(3844)
	Dictionary = []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
		"k", "l", "m", "n", "o", "p", "q", "r", "s", "t",
		"u", "v", "w", "x", "y", "z", "A", "B", "C", "D",
		"E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X",
		"Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7",
		"8", "9",
	}
)

func (u *URLCodec) Encode(URI string) int64 {
	return u.codec.DicToDec(URI) - u.startValue
}

func (u *URLCodec) Decode(encoded int64) string {
	return u.codec.DecToDic(encoded + u.startValue)
}

func NewURLCodec(startValue int64, dictionary []string) *URLCodec {
	return &URLCodec{
		startValue: startValue,
		codec:      NewCodec(dictionary),
	}
}

func New() *URLCodec {
	return NewURLCodec(StartValue, Dictionary)
}
