package router

import (
	_ "getcharzp.cn/docs"
	"getcharzp.cn/middlewares"
	"getcharzp.cn/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())

	// Swagger 配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 公共方法
	// 问题列表
	r.GET("/problem-list", service.GetProblemList)
	// 问题详情
	r.GET("/problem-detail", service.GetProblemDetail)
	// 用户
	r.GET("/user-detail", service.GetUserDetail)
	// 登录
	r.POST("/login", service.Login)
	// 发送验证码
	r.POST("/send-code", service.SendCode)
	// 注册
	r.POST("/register", service.Register)
	// 排行榜
	r.GET("/rank-list", service.GetRankList)
	// 提交记录
	r.GET("/submit-list", service.GetSubmitList)
	// 分类列表
	r.GET("/category-list", service.GetCategoryList)
	// 题解列表
	r.GET("/explanation-list", service.ExplanationList)
	// 题解详情
	r.GET("/explanation-detail", service.ExplanationDetail)
	// 题解阅读量
	r.GET("/explanation-readNum", service.ExplanationReadNum)
	// 评论列表
	r.GET("/comment", service.Comment)
	// 回复列表
	r.GET("/reply", service.Reply)

	// 管理员私有方法
	authAdmin := r.Group("/admin", middlewares.AuthAdminCheck())
	// 问题创建
	authAdmin.POST("/problem-create", service.ProblemCreate)
	// 问题修改
	authAdmin.PUT("/problem-modify", service.ProblemModify)
	// 分类创建
	authAdmin.POST("/category-create", service.CategoryCreate)
	// 分类修改
	authAdmin.PUT("/category-modify", service.CategoryModify)
	// 分类删除
	authAdmin.DELETE("/category-delete", service.CategoryDelete)
	// 获取测试案例
	authAdmin.GET("/test-case", service.GetTestCase)

	// 用户私有方法
	authUser := r.Group("/user", middlewares.AuthUserCheck())
	// 代码提交
	authUser.POST("/submit", service.Submit)
	// 问题收藏
	authUser.GET("/add-collect", service.AddCollect)
	// 收藏列表
	authUser.GET("/collect", service.GetCollectList)
	// 取消收藏
	authUser.DELETE("/collect-delete", service.CollectDelete)
	// 创建题解
	authUser.POST("/explanation-create", service.ExplanationCreate)
	// 修改题解
	authUser.PUT("/explanation-modify", service.ExplanationModify)
	// 删除题解
	authUser.DELETE("/explanation-delete", service.ExplanationDelete)
	// 题解点赞
	authUser.GET("/explanation-likeNum", service.ExplanationLikeNum)
	// 创建评论
	authUser.POST("/comment-create", service.CommentCreate)
	// 删除评论
	authUser.DELETE("/comment-delete", service.CommentDelete)
	// 创建回复
	authUser.POST("/reply-create", service.ReplyCreate)
	// 删除回复
	authUser.DELETE("/reply-delete", service.ReplyDelete)

	return r
}
