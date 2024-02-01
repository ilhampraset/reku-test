package service

import (
	"crypto/rand"
	"fmt"
	"reku-code-test/url-shortener/config"
	"reku-code-test/url-shortener/entity"
	"reku-code-test/url-shortener/repository"
	"strings"
)

type UrlService struct {
	repo *repository.MemoryStore
}

func NewURLService(repo *repository.MemoryStore) *UrlService {
	return &UrlService{repo: repo}
}

func (srv *UrlService) CreateShortURL(reqData *entity.Url) (string, error) {

	shortUrl := srv.generateShortURL(8)
	reqData.ShortURL = shortUrl
	err := srv.repo.Create(reqData)
	if err != nil {
		return "", err
	}
	return shortUrl, nil
}

func (srv *UrlService) GeByShortUrl(shortUrl string) (*entity.Url, error) {
	urls, err := srv.repo.GetByShort(shortUrl)
	if err != nil {
		return nil, err
	}
	return urls, nil
}

func (srv *UrlService) GetAllURLs(descending bool) ([]*entity.Url, error) {
	urls, err := srv.repo.GetAll(descending) // Set to true for descending order
	if err != nil {
		return nil, err
	}
	for i, u := range urls {
		if !strings.HasPrefix(u.ShortURL, config.HOST) {
			urls[i].ShortURL = fmt.Sprintf("%s/%s", config.HOST, u.ShortURL)
		}
	}
	return urls, nil
}

func (srv *UrlService) generateShortURL(length int) string {
	characterSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	characterSetLength := len(characterSet)

	randomBytes := make([]byte, length)
	rand.Read(randomBytes)

	var result string
	for _, b := range randomBytes {
		result += string(characterSet[int(b)%characterSetLength])
	}

	return result
}
