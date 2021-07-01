package main

import (
	"flag"
	"log"
)

func main() {
	var (
		meta string // 下载metadata缓存路径
		pkg  string // 下载指定API类目metadata
	)
	flag.StringVar(&meta, "meta", "", "metadata dir")
	flag.StringVar(&pkg, "pkg", "", "specific pkg")
	flag.Parse()
	if meta == "" {
		log.Fatalln("missing metadata dir")
	}
	err := Download(meta, pkg)
	if err != nil {
		log.Fatalln(err)
	}
}
