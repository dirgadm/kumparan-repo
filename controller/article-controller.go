package controller

import (
	"kumparan/cache"
	models "kumparan/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func FindArticles(c *gin.Context) {
	// var articles []models.Articles
	// models.DB.Find(&articles)

	// c.JSON(http.StatusOK, gin.H{"data": articles})

	title, _ := c.GetQuery("title")
	body, _ := c.GetQuery("body")
	author, _ := c.GetQuery("author")

	queryBody := "%" + body + "%"
	queryTitle := "%" + title + "%"

	mAuthor := make(map[string]interface{})
	mAuthor["author"] = author

	m := make(map[string]interface{})
	m["body"] = queryBody
	m["title"] = queryTitle

	var articles []models.Articles
	if author != "" {
		models.DB.Order("id ASC").Where(mAuthor).Where("title like ? and body like ? ", queryTitle, queryBody).Find(&articles)
	} else if title != "" || body != "" {
		models.DB.Order("id ASC").Where("title like ? and body like ? ", queryTitle, queryBody).Find(&articles)
	} else {
		models.DB.Order("id ASC").Find(&articles)
	}

	c.JSON(http.StatusOK, gin.H{"data": articles})
}

func CreateArticles(c *gin.Context) {
	var hc cache.ArticleCache
	hc = cache.NewRedisCache("127.0.0.1:6379", 1, 100)
	var articles models.Articles

	// Validate input
	if err := c.ShouldBindJSON(&articles); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert articles
	articles.Created = time.Now()
	models.DB.Create(&articles)

	hc.Set(strconv.Itoa(articles.ID), &articles)
	c.JSON(http.StatusOK, gin.H{"data": articles})
}

func QueryString(cond map[string]string) (queryString, param string) {

	// var where string
	// var param string

	// for k, v := range cond {
	// 	if cond[k] != ""{
	// 		where += k + "like " + v
	// 	}

	// }

	// if cond["body"] != ""{
	// 	where += "body like " + cond["body"]
	// }
	// if cond["param"] != ""{
	// 	where += "param like " + cond["param"]
	// }

	return "nil", "nil"
}
