/*
 * @Author:       SJX
 * @E-Mail:       540643428@qq.com
 * @TODO:         Become Better
 * @: ------------------------------------------------------
 * @Date: 2023-04-28 14:16:02
 */
package main

import (
	"github.com/cuizhu12138/HrbustDownLoadStation/controller"
	"github.com/cuizhu12138/HrbustDownLoadStation/service"
	"github.com/gin-gonic/gin"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")             // 允许所有来源进行跨域请求
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")          // 允许GET请求方法
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type") // 允许Content-Type请求头

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // 预检请求的响应，直接返回204 No Content
			return
		}

		c.Next()
	}
}

func InitRouter(r *gin.Engine) {

	r.StaticFS("/static", gin.Dir("cloud", true))

	r.Use(corsMiddleware())

	Normal := r.Group("/normal")
	Admin := r.Group("/admin")

	// 根路径重定向 到
	{

	}
	// normal api 用于学生日常下载和查看文件
	{
		Normal.GET("/list/", service.GetList)
		Normal.GET("/download/:filename", service.DownloadFile)
	}
	// admin api 用于管理员上传删除下载文件
	{
		Admin.GET("/login/", controller.Login)
	}
}
