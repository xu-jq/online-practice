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

// AddCollect
// @Tags 用户私有方法
// @Summary 添加收藏
// @Accept json
// @Param authorization header string true "authorization"
// @Param problem_identity query string true "problem_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/add-collect [get]
func AddCollect(c *gin.Context) {
	problemIdentity := c.Query("problem_identity")
	u, _ := c.Get("user_claims")
	userClaim := u.(*helper.UserClaims)
	userIdentity := userClaim.Identity
	if problemIdentity == "" || userIdentity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数错误",
		})
		return
	}
	var cnt int64
	err := models.DB.Where("problem_identity = ? AND user_identity = ?", problemIdentity, userIdentity).
		Model(new(models.ProblemCollect)).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Add Collect Error:" + err.Error(),
		})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "该问题已被用户收藏",
		})
		return
	}
	collect := &models.ProblemCollect{
		Identity:        helper.GetUUID(),
		ProblemIdentity: problemIdentity,
		UserIdentity:    userIdentity,
	}
	err = models.DB.Create(collect).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Create collect Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "收藏成功",
	})
	return
}

// GetCollectList
// @Tags 用户私有方法
// @Summary 收藏列表
// @Accept json
// @Param authorization header string true "authorization"
// @Param page query int false "page"
// @Param size query int false "size"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/collect [get]
func GetCollectList(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetProblemList Page strconv Error:", err)
		return
	}
	page = (page - 1) * size
	var count int64
	u, _ := c.Get("user_claims")
	userClaim := u.(*helper.UserClaims)
	userIdentity := userClaim.Identity
	list := make([]*models.ProblemCollect, 0)
	//获取用户收藏problemIdentity
	err = models.DB.Model(new(models.ProblemCollect)).Count(&count).
		Offset(page).Limit(size).Where("user_identity = ?", userIdentity).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get CollectProblemIdentity Error:" + err.Error(),
		})
		return
	}
	// 通过ProblemIdentity获取收藏问题列表
	CollectList := make([]*define.ProblemCollect, 0)
	CollectLists := make([]*define.ProblemCollect, 0)
	for _, collect := range list {
		problemIdentity := collect.ProblemIdentity
		err = models.DB.Model(new(models.ProblemBasic)).Where("identity = ?", problemIdentity).Find(&CollectList).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "Get CollectList Error:" + err.Error(),
			})
			return
		}
		CollectLists = append(CollectLists, CollectList...)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"count": count,
			"list":  CollectLists,
		},
	})
}

// CollectDelete
// @Tags 用户私有方法
// @Summary 取消收藏
// @Accept json
// @Param authorization header string true "authorization"
// @Param problem_identity query string true "problem_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/collect-delete [delete]
func CollectDelete(c *gin.Context) {
	u, _ := c.Get("user_claims")
	userClaim := u.(*helper.UserClaims)
	userIdentity := userClaim.Identity
	ProblemIdentity := c.Query("problem_identity")
	err := models.DB.Where("user_identity = ? AND problem_identity = ?", userIdentity, ProblemIdentity).Delete(new(models.ProblemCollect)).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Delete Collect Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}
