package middleware

import (
	"file-manager/svc"
	"strings"

	"github.com/gin-gonic/gin"
)

type ShortRouteMiddleware struct {
	ctx *svc.ServiceContext
}

func NewShortRouteMiddleware(ctx *svc.ServiceContext) *ShortRouteMiddleware {
	return &ShortRouteMiddleware{ctx}
}

func (s *ShortRouteMiddleware) validateShortRoute(shortRoute string) bool {
	exists := s.ctx.Cache.Exists(shortRoute)
	if !exists {
		return false
	}
	s.ctx.Cache.Delete(shortRoute)
	return true
}

func (s *ShortRouteMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		shortRoute := strings.Trim(path, "/")
		var resourcePath string
		exist := s.ctx.Cache.Get(shortRoute, &resourcePath)
		if !exist {
			c.AbortWithStatusJSON(404, gin.H{"error": "resource not found"})
			return
		}
		// 设置响应头，告知浏览器这是一个文件下载请求
		c.Header("Content-Disposition", "attachment; filename="+resourcePath)
		c.Header("Content-Type", "application/octet-stream")
		// 发送文件给客户端
		c.File(resourcePath)
		s.ctx.Cache.Delete(shortRoute)
		c.Next()
	}
}
