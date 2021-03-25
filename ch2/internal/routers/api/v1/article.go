package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {
	//test HTTP interface
	// app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	c.JSON(http.StatusOK, gin.H{})
	return
}
func (a Article) List(c *gin.Context)   {}
func (a Article) Create(c *gin.Context) {}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}
