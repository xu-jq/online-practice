package service

import (
	"getcharzp.cn/define"
	"getcharzp.cn/helper"
	"getcharzp.cn/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// CommentCreate
// @Tags 用户私有方法
// @Summary 评论创建
// @Accept json
// @Param authorization header string true "authorization"
// @Param data body define.CommentCreate true "CommentCreate"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/comment-create [post]
func CommentCreate(c *gin.Context) {
	in := new(define.CommentCreate)
	err := c.ShouldBindJSON(in)
	if err != nil {
		log.Println("[JsonBind Error] : ", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数错误",
		})
		return
	}
	u, _ := c.Get("user_claims")
	userClaim := u.(*helper.UserClaims)
	userIdentity := userClaim.Identity
	comment := &models.Comment{
		Identity:        helper.GetUUID(),
		ContentIdentity: in.ContentIdentity,
		UserIdentity:    userIdentity,
		Comment:         in.Comment,
	}
	err = models.DB.Create(comment).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Create Comment Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "评论成功",
	})
}

// Comment
// @Tags 公共方法
// @Summary 评论列表
// @Accept json
// @Param content_identity query string true "content_identity"
// @Param page query int false "page"
// @Param size query int false "size"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /comment [get]
func Comment(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetProblemList Page strconv Error:", err)
		return
	}
	page = (page - 1) * size
	var count int64
	contentIdentity := c.Query("content_identity")
	list := make([]*define.CommentList, 0)
	err = models.DB.Model(new(models.Comment)).Where("content_identity = ?", contentIdentity).Count(&count).
		Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Comment Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"count": count,
			"list":  list,
		},
	})
}

// CommentDelete
// @Tags 用户私有方法
// @Summary 评论删除
// @Accept json
// @Param authorization header string true "authorization"
// @Param user_identity query string true "user_identity"
// @Param comment_identity query string true "comment_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/comment-delete [delete]
func CommentDelete(c *gin.Context) {
	user := c.Query("user_identity")
	identity := c.Query("comment_identity")
	u, _ := c.Get("user_claims")
	userClaim := u.(*helper.UserClaims)
	userIdentity := userClaim.Identity
	if userIdentity != user {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户信息不匹配，无法删除",
		})
		return
	}
	err := models.DB.Where("identity = ?", identity).Delete(new(models.Comment)).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Delete Comment err:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}
