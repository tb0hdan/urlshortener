package memory

import (
	"fmt"
	"log"
	"os"

	"github.com/tb0hdan/memcache"
)

// Logger - log.Logger extension
type Logger struct {
	st *log.Logger
}

// Printf - same as log.Printf
func (l *Logger) Printf(fmt string, args ...interface{}) {
	l.st.Printf(fmt, args...)
}

// Debug - same as log.Println
func (l *Logger) Debug(s ...interface{}) {
	l.st.Println(s...)
}

// NewLog - Create new logger instance
func NewLog() *Logger {
	return &Logger{st: log.New(os.Stderr, "", log.LstdFlags)}
}

// Storage - URL storage based on memcache.CacheType
type Storage struct {
	cache *memcache.CacheType
}

// GetByID - return URL value by its integer id
func (m *Storage) GetByID(urlID int64) (url string) {
	return fmt.Sprintf("%s", m.cache.GetByID(urlID))
}

// GetByLongURL - check whether long URL is in storage
func (m *Storage) GetByLongURL(long string) (value string, ok bool) {
	val, ok := m.cache.Get(long)
	if ok {
		value = val.(string)
	}

	return
}

// Add - Add shortened and original URL to storage and return its id
func (m *Storage) Add(short, long string) (urlID int64) {
	return m.cache.Add(long, short)
}

// Len - Return storage length (unsafe for concurrent access)
func (m *Storage) Len() (storageLen int64) {
	return m.cache.Len()
}

// LenSafe - Return storage length
func (m *Storage) LenSafe() (storageLen int64) {
	return m.cache.LenSafe()
}

// NewStorage - New memory storage instance
func NewStorage() *Storage {
	return &Storage{
		cache: memcache.New(NewLog()),
	}
}
