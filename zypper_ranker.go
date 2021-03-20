package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const (
	downloadAppName     = "nano"
	tumbleweedUrlSuffix = "tumbleweed/repo/oss/"
)

var (
	zypperPrefix string
)

func main() {

	zypperInit()

	csv_mirror, _ := os.Open("./mirrors.csv")
	mirrorlists := csv.NewReader(csv_mirror)

	for {
		mirror, err := mirrorlists.Read()
		// mirror[0] => name
		//       [1] => url without "tumbleweed/repo/oss/"
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("csv reading failed")
		}

		zypperForceSingleRepo(mirror[1] + tumbleweedUrlSuffix)

		// Refresh test
		beginTime := time.Now()
		callZypper("refresh")
		duration := time.Since(beginTime).String()
		fmt.Println("Refresh test: " + mirror[0] + " duration:" + duration)

		// installation test
		beginTime = time.Now()
		callZypper("install --no-recommends --no-confirm --download-only " + downloadAppName)
		duration = time.Since(beginTime).String()
		fmt.Println("Install test: " + mirror[0] + " duration:" + duration)
	}
}

func zypperInit() {
	cacheDir, e := os.UserCacheDir()
	check(e, "~./.cache/ is inaccessible")
	zypperPrefix = "sudo zypper --no-gpg-checks --root=" + filepath.Join(cacheDir, "tumbleweed") + " "
	callZypper("-v")
}

func callZypper(command string) {

	zypperCompleteCmd := zypperPrefix + command
	fmt.Println("[->] " + zypperCompleteCmd)
	_, err := exec.Command("/bin/sh", "-c", zypperCompleteCmd).Output()
	check(err, "Failed: "+zypperCompleteCmd)
}

func zypperForceSingleRepo(url string) {
	callZypper("removerepo -a")
	callZypper("ar -c " + url + " Tumbleweed")
}

func check(e error, s string) {
	if e != nil {
		println(s)
		panic(e)
	}
}
