package utils

import (
	"crypto/md5"
	"encoding/hex"
	"file-manager/svc"
	"fmt"
	"math/rand"
	"time"
)

func GenerateShortRoute(filePath string, ctx *svc.ServiceContext) string {
	// 获取当前时间戳
	timestamp := time.Now().Unix()
	token := RandomString(10)
	// 合并时间戳和token
	combined := fmt.Sprintf("%d%s", timestamp, token)
	// 计算MD5哈希值
	hash := md5.Sum([]byte(combined))
	shortRoute := hex.EncodeToString(hash[:])
	ctx.Cache.Set(shortRoute, filePath, 10*time.Minute)
	return shortRoute
}

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
