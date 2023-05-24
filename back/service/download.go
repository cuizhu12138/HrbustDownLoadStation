/*
 * @Author:       SJX
 * @E-Mail:       540643428@qq.com
 * @TODO:         Become Better
 * @: ------------------------------------------------------
 * @Date: 2023-05-24 09:56:35
 */
package service

import (
	"fmt"

	"github.com/cuizhu12138/HrbustDownLoadStation/etc"
	"github.com/gin-gonic/gin"
)

func DownloadFile(c *gin.Context) {
	filename := c.Param("filename")

	FullPath := etc.FilePath + "/" + filename

	fmt.Println("debug :" + FullPath)

	// 设置文件下载头部
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")

	c.File(FullPath)
}
