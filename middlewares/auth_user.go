package middlewares

import (
	"getcharzp.cn/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthUserCheck() gin.HandlerFunc {
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6ImZhOTNmODg0LTJjNzQtNDlmYy1hMjI1LWJhNDAyMDFmNGUxMiIsIm5hbWUiOiJ1c2VyMSIsImlzX2FkbWluIjowfQ.RTRSo5V5U-5vPDx-CoIY9vj7Ffx-AqsGnp6_aA85YWg
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		userClaim, err := helper.AnalyseToken(auth)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Unauthorized Authorization",
			})
			return
		}
		if userClaim == nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Unauthorized Admin",
			})
			return
		}
		c.Set("user_claims", userClaim)
		c.Next()
	}
}
