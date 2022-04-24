package main

import (
	"fmt"
	"github.com/kininaru/mc-server-backup/upload"
	zip_tools "github.com/kininaru/mc-server-backup/zip-tools"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 5 {
		log.Fatalln("Arguments not enough")
	}
	if _, err := os.Stat("./mc-backup"); os.IsNotExist(err) {
		_ = os.Mkdir("./mc-backup", 0777)
	}
	filename := "backup" + strconv.Itoa(int(time.Now().UnixNano())) + ".zip"
	zip_tools.ToZip("./mc", "./mc-backup/"+filename)
	fmt.Println("Zipping finished. Uploading...")
	upload.ToAliyun(os.Args[1], os.Args[2], os.Args[3], os.Args[4], filename)
	fmt.Println("Finished. Filename is " + filename)
}
