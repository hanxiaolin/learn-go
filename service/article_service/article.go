package article_service

import (
	"hanxiaolin/gin-demo/logging"
	"hanxiaolin/gin-demo/models"
	"hanxiaolin/gin-demo/pkg/gredis"
	"hanxiaolin/gin-demo/service/cache_service"
	"encoding/json"
	"fmt"
)

type Article struct {
	ID            int
	TagID         int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         int
	CreatedBy     string
	ModifiedBy    string

	PageNum  int
	PageSize int
}

func (a *Article) ExistByID() (bool, error) {
	return models.ExistArticleByID(a.ID)
}

func (a *Article) Get() (*models.Article, error) {
	var cacheArticle *models.Article

	cache := cache_service.Article{ID: a.ID}
	key := cache.GetArticleKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheArticle)
			return cacheArticle, nil
		}
	}

	article, err := models.GetArticle(a.ID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = gredis.Set(key, article, 3600)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return article, nil
}
