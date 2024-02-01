package repository

import (
	"net/url"
	"reku-code-test/url-shortener/entity"
	"sort"
	"sync"
)

type MemoryStore struct {
	urls    map[string]*entity.Url
	counter int
	mu      sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		urls: make(map[string]*entity.Url),
	}
}

func (repo *MemoryStore) Create(urlData *entity.Url) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.urls[urlData.ShortURL] = urlData

	return nil

}

func (repo *MemoryStore) GetAll(descending bool) ([]*entity.Url, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	var urls []*entity.Url
	for _, urlData := range repo.urls {
		urls = append(urls, urlData)
	}

	sort.Slice(urls, func(i, j int) bool {
		if descending {
			return urls[i].ClickCount > urls[j].ClickCount
		}
		return urls[i].ClickCount < urls[j].ClickCount
	})

	return urls, nil
}

func (repo *MemoryStore) GetByShort(shortUrl string) (*entity.Url, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	decodedShortUrl, err := url.PathUnescape(shortUrl)
	if err != nil {
		return nil, err
	}
	return repo.urls[decodedShortUrl], nil
}
