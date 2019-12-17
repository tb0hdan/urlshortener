package storage

// GenericStorage - Definition for URL storage
type GenericStorage interface {
	// GetByID - return URL value by its integer id
	GetByID(urlID int64) (url string)
	// GetByLongURL - check whether long URL is in storage
	GetByLongURL(long string) (value string, ok bool)
	// Add - Add shortened and original URL to storage and return its id
	Add(short, long string) (urlID int64)
	// Len - Return storage length (unsafe for concurrent access)
	Len() (storageLen int64)
	// Len - Return storage length
	LenSafe() (storageLen int64)
}
