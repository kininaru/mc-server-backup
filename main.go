package main

import (
	"fmt"
	zip_tools "github.com/kininaru/mc-server-backup/zip-tools"
	"os"
	"strconv"
	"time"
)

func main() {
	if _, err := os.Stat("./mc-backup"); os.IsNotExist(err) {
		_ = os.Mkdir("./mc-backup", 0777)
	}
	zip_tools.ToZip("./test", "./mc-backup/backup"+strconv.Itoa(int(time.Now().UnixNano()))+".zip")
	fmt.Println("Zipping finished. Uploading...")
}
