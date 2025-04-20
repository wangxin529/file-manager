package controller

import (
	"file-manager/svc"
	"file-manager/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type FileOperator struct {
	ctx *svc.ServiceContext
}

func NewFileOperator(ctx *svc.ServiceContext) *FileOperator {
	return &FileOperator{ctx: ctx}
}

// func (f *FileOperator) GetResource(c *gin.Context) {
// 	// TODO: 实现GetResource方法
// 	c.JSON(http.StatusNotImplemented, gin.H{"error": "方法未实现"})
// }

func (f *FileOperator) GetResourceID(c *gin.Context) {
	filename := c.Param("filepath")
	filename = filepath.Join(f.ctx.Config.Application.StoragePath, filename)
	route := utils.GenerateShortRoute(filename, f.ctx)
	c.JSON(http.StatusOK, gin.H{
		"route": route,
	})
}

func (f *FileOperator) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	filename := filepath.Join(f.ctx.Config.Application.StoragePath, file.Filename)
	err = c.SaveUploadedFile(file, filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shortRouter := utils.GenerateShortRoute(filename, f.ctx)
	c.JSON(http.StatusOK, gin.H{
		"message": "文件上传成功",
		"path":    shortRouter,
	})
}
