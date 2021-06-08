package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func responseError(j *JsonResponse, msg string) {
	j.Code = -1
	j.Msg = msg
}

func (c *DocController) Post() {
	var params mdInfs
	jsonRes := &JsonResponse{
		Code: 0,
		Msg:  "写入成功",
	}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	if err != nil {
		responseError(jsonRes, "无数据上传")
		c.Ctx.ResponseWriter.Status = 301
		fmt.Println(err.Error())
	}

	filePath := strings.Join([]string{params.Filename, time.Now().Format("15:04:05")}, "-")
	basePath := strings.Join([]string{"mdfiles", params.Authorname}, "/")
	dirPath := strings.Join([]string{basePath, time.Now().Format("2006-01-02")}, "/")
	compPath := strings.Join([]string{dirPath, filePath}, "/")
	latestPath := strings.Join([]string{basePath, "latestMD.md"}, "/")

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			fmt.Println("目录创建错误, 原因：", err.Error())
		}
	}
	// f, _ := os.Create(compPath)
	// num, err := f.WriteString(params.Content)
	fmt.Println(compPath, latestPath)
	contentByte := []byte(params.Content)
	errWrite := ioutil.WriteFile(compPath, contentByte, os.ModePerm)
	if errWrite != nil {
		fmt.Println("文件写入错误, 原因：", errWrite.Error())
		responseError(jsonRes, "写入失败")
		c.Ctx.ResponseWriter.Status = 500
	}

	errWrite = ioutil.WriteFile(latestPath, contentByte, os.ModePerm)
	if errWrite != nil {
		fmt.Println("文件写入错误, 原因：", errWrite.Error())
		responseError(jsonRes, "写入失败")
		c.Ctx.ResponseWriter.Status = 500
	}

	c.Data["json"] = jsonRes
	_ = c.ServeJSON()
}

func (c *DocController) Get() {
	jsonRes := &JsonResponse{
		Code: 0,
		Msg:  "传输成功！",
	}
	params := &mdInfs{
		Authorname: c.GetString("authorName"),
	}
	dirPath := strings.Join([]string{"mdfiles", params.Authorname, "latestMD.md"}, "/")
	fmt.Println(params, dirPath)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		responseError(jsonRes, "未创建任何文档！")
	} else {
		byteData, err := ioutil.ReadFile(dirPath)
		if err != nil {
			responseError(jsonRes, "打开文件失败！")
			fmt.Println("打开文件失败: ", err)
		} else {
			jsonRes.Msg = string(byteData[:])
		}
	}

	c.Data["json"] = jsonRes
	fmt.Println(c.Data["json"])
	_ = c.ServeJSON()

}
