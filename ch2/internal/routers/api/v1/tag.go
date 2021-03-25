package v1

import "github.com/gin-gonic/gin"

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {}

// @Summary Getting list of tags
// @Produce  json
// @Param name query string false "tag name" maxlength(100)
// @Param state query int false "status" Enums(0, 1) default(1)
// @Param page query int false "page num"
// @Param page_size query int false "page display num"
// @Success 200 {object} errcode.Error "success"
// @Failure 400 {object} errcode.Error "Bad request"
// @Failure 500 {object} errcode.Error "Internal Server Error"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context)   {}
func (t Tag) Create(c *gin.Context) {}
func (t Tag) Update(c *gin.Context) {}
func (t Tag) Delete(c *gin.Context) {}
