package urlstorage

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

func (us *URLStorage) GetByLongURL(long string) (value string, ok bool) {
	us.m.RLock()
	defer us.m.RUnlock()
	value, ok = us.URLsHash[long]

	return
}

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

func (us *URLStorage) Len() (storageLen int64) {
	storageLen = int64(len(us.URLs))
	return
}

func (us *URLStorage) LenSafe() (storageLen int64) {
	us.m.RLock()
	defer us.m.RUnlock()
	storageLen = us.Len()

	return
}
