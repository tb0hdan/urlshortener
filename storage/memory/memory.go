package memory

// GetByID - return URL value by its integer id
func (us *URLStorage) GetByID(urlID int64) (url string) {
	us.m.RLock()
	defer us.m.RUnlock()

	if urlID < 0 {
		return ""
	}

	if urlID <= us.Len() {
		urlItem := us.URLs[urlID-1]
		url = urlItem.Long
	}

	return
}

// GetByLongURL - check whether long URL is in storage
func (us *URLStorage) GetByLongURL(long string) (value string, ok bool) {
	us.m.RLock()
	defer us.m.RUnlock()
	value, ok = us.URLsHash[long]

	return
}

// Add - Add shortened and original URL to storage and return its id
func (us *URLStorage) Add(short, long string) (urlID int64) {
	if _, ok := us.GetByLongURL(long); ok {
		return 0
	}

	us.m.Lock()
	defer us.m.Unlock()

	us.URLs = append(us.URLs, &URLItem{
		Short: short,
		Long:  long,
	})
	us.URLsHash[long] = short

	return us.Len()
}

// Len - Return storage length (unsafe for concurrent access)
func (us *URLStorage) Len() (storageLen int64) {
	storageLen = int64(len(us.URLs))
	return
}

// LenSafe - Return storage length
func (us *URLStorage) LenSafe() (storageLen int64) {
	us.m.RLock()
	defer us.m.RUnlock()
	storageLen = us.Len()

	return
}
