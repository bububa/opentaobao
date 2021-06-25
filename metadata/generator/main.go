package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

var (
	GitTag      string // set at compile time with -ldflags
	GitRevision string // set at compile time with -ldflags
	GitSummary  string // set at compile time with -ldflags
)

func main() {
	var (
		meta  string
		patch string
		pkg   string
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
	tpl := map[string]string{
		"GitTag":      GitTag,
		"GitRevision": GitRevision,
		"GitSummary":  GitSummary,
	}
	tmpl.Execute(fd, tpl)
	return nil
}
