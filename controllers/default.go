package controllers

import (
	"fmt"
	"os"
	"strings"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

type DocController struct {
	beego.Controller
}

type JsonResponse struct {
	Code int
	Msg  string
}

type mdInfs struct {
	authorName string
	filename   string
	content    string
}

func (c *DocController) Post() {
	jsonRes := &JsonResponse{
		Code: 200,
		Msg:  "写入成功",
	}
	params := &mdInfs{
		authorName: c.GetString("authorName"),
		filename:   c.GetString("filename"),
		content:    c.GetString("content"),
	}

	filePath := strings.Join([]string{params.filename, time.Now().Format("15:04:05")}, "-")
	dirPath := strings.Join([]string{"mdfiles", params.authorName, time.Now().Format("2006-01-02")}, "/")
	compPath := dirPath + "/" + filePath

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0644)
		if err != nil {
			fmt.Println("目录创建错误, 原因：", err)
		}
	}
	f, _ := os.Create(compPath)
	_, err := f.WriteString(params.content)

	if err != nil {
		fmt.Println("文件写入错误, 原因：", err)
		jsonRes.Code = 500
		jsonRes.Msg = "写入失败"
	}

	c.Data["json"] = jsonRes
	_ = c.ServeJSON()
}
