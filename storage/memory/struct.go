package memory

import "sync"

// URLItem - Relevant URL information, i.e. short and original (long) links
type URLItem struct {
	Short string
	Long  string
}

// URLStorage - URL Storage with bound methods
type URLStorage struct {
	URLs     []*URLItem
	URLsHash map[string]string
	m        sync.RWMutex
}
