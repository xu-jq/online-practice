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

// ReplyCreate
// @Tags 用户私有方法
// @Summary 回复创建
// @Accept json
// @Param authorization header string true "authorization"
// @Param data body define.ReplyCreate true "ReplyCreate"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/reply-create [post]
func ReplyCreate(c *gin.Context) {
	in := new(define.ReplyCreate)
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
	reply := &models.Reply{
		Identity:        helper.GetUUID(),
		CommentIdentity: in.CommentIdentity,
		UserIdentity:    userIdentity,
		Reply:           in.Reply,
	}
	err = models.DB.Create(reply).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Create Reply Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "回复成功",
	})
}

// Reply
// @Tags 公共方法
// @Summary 回复列表
// @Accept json
// @Param comment_identity query string true "comment_identity"
// @Param page query int false "page"
// @Param size query int false "size"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /reply [get]
func Reply(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetList Page strconv Error:", err)
		return
	}
	page = (page - 1) * size
	var count int64
	commentIdentity := c.Query("comment_identity")
	list := make([]*define.ReplyList, 0)
	err = models.DB.Model(new(models.Reply)).Where("comment_identity = ?", commentIdentity).Count(&count).
		Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Reply Error:" + err.Error(),
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

// ReplyDelete
// @Tags 用户私有方法
// @Summary 回复删除
// @Accept json
// @Param authorization header string true "authorization"
// @Param user_identity query string true "user_identity"
// @Param reply_identity query string true "reply_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/reply-delete [delete]
func ReplyDelete(c *gin.Context) {
	user := c.Query("user_identity")
	identity := c.Query("reply_identity")
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
	err := models.DB.Where("identity = ?", identity).Delete(new(models.Reply)).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Delete Reply err:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}
