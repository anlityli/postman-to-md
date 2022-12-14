package main

import (
	"github.com/gogf/gf/os/glog"
	"postman-to-md/app/service"
)

func main() {
	if err := service.PmReader.Run(); err != nil {
		glog.Fatal(err)
	}
}
