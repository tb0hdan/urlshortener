package urlstorage

import "sync"

type URLItem struct {
	Short string
	Long  string
}

type URLStorage struct {
	URLs     []*URLItem
	URLsHash map[string]string
	m        sync.RWMutex
}
