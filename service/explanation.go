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

// ExplanationList
// @Tags 公共方法
// @Summary 题解列表
// @Accept json
// @Param problem_identity query string true "problem_identity"
// @Param page query int false "page"
// @Param size query int false "size"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /explanation-list [get]
func ExplanationList(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetProblemList Page strconv Error:", err)
		return
	}
	page = (page - 1) * size
	var count int64
	problemIdentity := c.Query("problem_identity")
	list := make([]*define.ExplanationList, 0)
	err = models.DB.Model(new(models.ProblemExplanation)).Where("problem_identity = ?", problemIdentity).Count(&count).
		Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get ExplanationList Error:" + err.Error(),
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

// ExplanationDetail
// @Tags 公共方法
// @Summary 题解详情
// @Accept json
// @Param explanation_identity query string true "explanation_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /explanation-detail [get]
func ExplanationDetail(c *gin.Context) {
	explanationIdentity := c.Query("explanation_identity")
	data := new(models.ProblemExplanation)
	err := models.DB.Where("identity = ?", explanationIdentity).First(data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get ExplanationDetail Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
	return
}

// ExplanationCreate
// @Tags 用户私有方法
// @Summary 题解创建
// @Accept json
// @Param authorization header string true "authorization"
// @Param data body define.ExplanationCreate true "ExplanationCreate"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/explanation-create [post]
func ExplanationCreate(c *gin.Context) {
	in := new(define.ExplanationCreate)
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
	explanation := &models.ProblemExplanation{
		Identity:        helper.GetUUID(),
		ProblemIdentity: in.ProblemIdentity,
		UserIdentity:    userIdentity,
		Title:           in.Title,
		Content:         in.Content,
		ReadNum:         0,
		LikeNum:         0,
	}
	err = models.DB.Create(explanation).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Create ProblemExplanation Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "题解创建成功",
	})
	return
}

// ExplanationModify
// @Tags 用户私有方法
// @Summary 题解修改
// @Accept json
// @Param authorization header string true "authorization"
// @Param data body define.ExplanationModify true "ExplanationModify"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/explanation-modify [put]
func ExplanationModify(c *gin.Context) {
	in := new(define.ExplanationModify)
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
	if userIdentity != in.UserIdentity {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户信息不匹配，无法更改",
		})
		return
	}
	explanation := &models.ProblemExplanation{
		Title:   in.Title,
		Content: in.Content,
	}
	err = models.DB.Where("identity = ?", in.Identity).Updates(explanation).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Modify Explanation err:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
}

// ExplanationDelete
// @Tags 用户私有方法
// @Summary 题解删除
// @Accept json
// @Param authorization header string true "authorization"
// @Param user_identity query string true "user_identity"
// @Param explanation_identity query string true "explanation_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/explanation-delete [delete]
func ExplanationDelete(c *gin.Context) {
	user := c.Query("user_identity")
	identity := c.Query("explanation_identity")
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
	err := models.DB.Where("identity = ?", identity).Delete(new(models.ProblemExplanation)).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Delete Explanation err:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

// ExplanationReadNum
// @Tags 公共方法
// @Summary 题解阅读量
// @Accept json
// @Param explanation_identity query string true "explanation_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /explanation-readNum [get]
func ExplanationReadNum(c *gin.Context) {
	identity := c.Query("explanation_identity")
	explanation := &models.ProblemExplanation{}
	err := models.DB.Where("identity = ?", identity).Select("read_num").First(&explanation).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Select ReadNum err:" + err.Error(),
		})
		return
	}
	explanation.ReadNum = explanation.ReadNum + 1
	err = models.DB.Where("identity = ?", identity).Updates(explanation).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Updates ReadNum err:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "阅读量+1成功",
	})
}

// ExplanationLikeNum
// @Tags 用户私有方法
// @Summary 题解点赞
// @Accept json
// @Param authorization header string true "authorization"
// @Param explanation_identity query string true "explanation_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/explanation-likeNum [get]
func ExplanationLikeNum(c *gin.Context) {
	identity := c.Query("explanation_identity")
	explanation := &models.ProblemExplanation{}
	err := models.DB.Where("identity = ?", identity).Select("like_num").First(explanation).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Select LikeNum err:" + err.Error(),
		})
		return
	}
	explanation.LikeNum = explanation.LikeNum + 1
	err = models.DB.Where("identity = ?", identity).Updates(explanation).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Updates LikeNum err:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "点赞成功",
	})
}
