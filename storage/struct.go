package storage

type GenericStorage interface {
	GetByID(urlID int64) (url string)
	GetByLongURL(long string) (value string, ok bool)
	Add(short, long string) (urlID int64)
	Len() (storageLen int64)
	LenSafe() (storageLen int64)
}
