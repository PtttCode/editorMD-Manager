package controllers

import (
	"encoding/json"
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
	Authorname string
	Filename   string
	Content    string
}

func (c *DocController) Post() {
	var params mdInfs
	jsonRes := &JsonResponse{
		Code: 0,
		Msg:  "写入成功",
	}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	if err != nil {
		jsonRes.Code = -1
		jsonRes.Msg = "无数据上传"
		c.Ctx.ResponseWriter.Status = 301
		fmt.Println(err.Error())
	}

	filePath := strings.Join([]string{params.Filename, time.Now().Format("15:04:05")}, "-")
	dirPath := strings.Join([]string{"mdfiles", params.Authorname, time.Now().Format("2006-01-02")}, "/")
	compPath := dirPath + "/" + filePath

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			fmt.Println("目录创建错误, 原因：", err.Error())
		}
	}
	f, _ := os.Create(compPath)
	num, err := f.WriteString(params.Content)
	fmt.Println(num)
	fmt.Println(params)

	if err != nil {
		fmt.Println("文件写入错误, 原因：", err.Error())
		jsonRes.Code = -1
		jsonRes.Msg = "写入失败"
		c.Ctx.ResponseWriter.Status = 500
	}

	c.Data["json"] = jsonRes
	_ = c.ServeJSON()
}
