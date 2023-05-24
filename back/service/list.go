/*
 * @Author:       SJX
 * @E-Mail:       540643428@qq.com
 * @TODO:         Become Better
 * @: ------------------------------------------------------
 * @Date: 2023-04-28 14:23:27
 */
package service

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"log"

	"github.com/cuizhu12138/HrbustDownLoadStation/midlleware"

	"github.com/cuizhu12138/HrbustDownLoadStation/etc"
)

func GetList(c *gin.Context) {
	// 设置日志前缀
	log.SetPrefix("List Request")

	// fileList := []string{"file1.txt", "file2.jpg", "file3.pdf"}
	var FileList []string

	// 要遍历的文件夹路径
	folderPath := etc.FilePath

	// 定义一个匿名函数作为 filepath.Walk 的回调函数
	walkFunc := func(path string, info os.FileInfo, err error) error {
		// 如果遍历到的是文件，则输出文件名
		if !info.IsDir() {
			FileList = append(FileList, info.Name())
		}
		return nil
	}

	// 遍历文件夹及其子文件夹，并执行回调函数
	err := filepath.Walk(folderPath, walkFunc)

	if err != nil {
		log.Println("Bad ListRequest")
		c.JSON(http.StatusBadRequest, nil)
	}

	// 创建包含文件列表的JSON对象
	response := midlleware.ResponseList{
		Files: FileList,
	}

	c.JSON(http.StatusOK, response)
}
