package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
	"time"
)

var (
	version = "dev"     // set at compile time with -ldflags
	commit  = "none"    // set at compile time with -ldflags
	date    = "unknown" // set at compile time with -ldflags
	GitDate = ""
	builtBy = "unknown" // set at compile time with -ldflags
)

func main() {
	var (
		meta  string // API文档metadata保存路径
		patch string // patch文件保存路径
		pkg   string // 生成指定API类目包
	)
	flag.StringVar(&meta, "meta", "", "metadata dir")
	flag.StringVar(&patch, "patch", "", "metadata patch dir")
	flag.StringVar(&pkg, "pkg", "", "specific pkg")
	flag.Parse()
	if meta == "" {
		log.Fatalln("missing metadata dir")
	}
	if err := Gen(meta, patch, pkg); err != nil {
		log.Fatalln(err)
	}
	if err := genGoDoc(); err != nil {
		log.Fatalln(err)
	}
}

// 生成根目录doc.go
func genGoDoc() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	filePath := filepath.Join(wd, "metadata/template/doc.tpl")
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		return err
	}
	targetFile := filepath.Join(wd, "doc.go")
	fd, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer fd.Close()
	if date == "unknown" && GitDate != "" {
		if t, err := strconv.ParseInt(GitDate, 10, 64); err == nil {
			date = time.Unix(t, 0).UTC().Format(time.RFC3339)
		}
	}
	tpl := map[string]string{
		"GitVersion": version,
		"GitCommit":  commit,
		"GitDate":    date,
		"GitBuiltBy": builtBy,
	}
	tmpl.Execute(fd, tpl)
	return nil
}
