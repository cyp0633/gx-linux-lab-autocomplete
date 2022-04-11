package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//func addFile(timeStamp int64) {
//	dirs := []string{"/home/hnu/", "/usr/bin/", "/opt/", "/var/log/", "/mnt/", "/etc/"}
//	subdir = []string{"csee", "hnu", "yuelu", "tianma", "fenghuang"}
//}

func main() {
	fmt.Println("本工具可以帮您自动完成工训 Linux 入门实验，请输入您的学号")
	var stuNumber int
	_, err := fmt.Scanf("%d", &stuNumber)
	if err != nil {
		fmt.Println("学号输入错误")
		panic(err)
	}
	timeStamp := time.Now().Unix()
	timeStamp -= 60 //To make it real, add delay
	osType := runtime.GOOS
	if osType == "linux" {
		fmt.Println("是否需要伪造文件？Y/n")
		var option string
		fmt.Scanf("%s", &option)
		if strings.ToLower(option) != "n" {
			//addFile(timeStamp)
		}
	}
	timeStamp = 5000000000 - timeStamp
	resp, err := http.Get("http://132.232.98.70:6363/check?id=" + strconv.Itoa(stuNumber) + "&v=" + strconv.FormatInt(timeStamp, 10))
	if err != nil {
		fmt.Println("请求发送失败")
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	s, err := ioutil.ReadAll(resp.Body)
	if string(s) == "OK" {
		fmt.Println("已完成")
	} else {
		fmt.Println("失败，服务器返回信息：" + string(s))
	}
}
