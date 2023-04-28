package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	SetupLogger()
	simpleHttpGET("www.baidu.com")
	simpleHttpGET("https://www.baidu.com")
}

func SetupLogger() {
	logFileLocation, err := os.OpenFile("/Users/hezefeng/GolandProjects/LearnZap/GoLogger/logs/test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	if err != nil {
		log.Printf(err.Error())
	}

	log.SetOutput(logFileLocation)
}

// 简单的http连接
func simpleHttpGET(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching url %s : %s", url, err.Error())
	} else {
		log.Printf("Status Code for %s : %s", url, resp.Status)
		err := resp.Body.Close()
		if err != nil {
			return
		}

	}
}
