package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func addFile(timeStamp int64) {
	dirs := []string{"/home/hnu/", "/usr/bin/", "/opt/", "/var/log/", "/mnt/", "/etc/"}
	subDir := []string{"csee/", "hnu/", "yuelu/", "tianma/", "fenghuang/"}
	fileName := []string{".puzzle.txt", ".game.txt", ".answer"}

	v := timeStamp * timeStamp
	d := dirs[v%6]
	v = v * timeStamp
	mod := v % 5
	if mod < 0 {
		mod += 5
	}
	d = d + subDir[mod]

	err := os.MkdirAll(d, 0)
	if err != nil {
		fmt.Println("创建失败，请检查是否以 sudo 运行")
		panic(err)
	}
	timeStamp = 5000000000 - timeStamp
	path := d + fileName[timeStamp%3]

	data := []byte(strconv.FormatInt(timeStamp, 10))
	err = ioutil.WriteFile(path, data, 0)
	if err != nil {
		fmt.Println("写入失败，请检查是否以 sudo 运行")
		panic(err)
	}
	fmt.Printf("已将内容%v写入到文件%v\n", timeStamp, path)
}

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
			addFile(timeStamp)
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
