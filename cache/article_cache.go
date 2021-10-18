package cache

import "kumparan/models"

type ArticleCache interface {
	Set(key string, value *models.Articles)
	Get(key string) *models.Articles
}
