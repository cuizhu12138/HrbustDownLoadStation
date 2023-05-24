/*
 * @Author:       SJX
 * @E-Mail:       540643428@qq.com
 * @TODO:         Become Better
 * @: ------------------------------------------------------
 * @Date: 2023-04-28 12:45:38
 */
package main

import (
	"log"
	"os"

	"github.com/cuizhu12138/HrbustDownLoadStation/etc"
	"github.com/gin-gonic/gin"
)

func main() {

	// 设置一下日志输出的路径
	logFile, err := os.OpenFile(etc.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}

	log.SetOutput(logFile)


	r := gin.Default()

	InitRouter(r)

	r.Run(":11451")
}

// 初始化一些服务
func InitService() {

}

/*
后端：
直接访问根目录("/"目录)，然后读取图标(列表通过静态资源浏览？再看一眼cloudlist的实现)

访问根目录的时候就应该初始化下载列表,所以根目录应该重定向到 ("/list")

list 接口应该返回一个结构体，包含文件名和对应的下载链接

网页上的列表也应该全部都是超链接，每个文件名指向下载链接

管理员上传就正常上传，用户通过刷新获得最新列表

管理员删除文件或者管理文件时可能需要先对文件加锁，然后中断所有正在下载的用户(所以这里文件分片发会好一些，每一片传送之前
都先判断下载链接是否应该关闭)

文件上传MD5校验


前端：
不能自动刷新！！可能会中断下载


*/
