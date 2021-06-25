package main

import (
	"flag"
	"log"
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
	err := Gen(meta, patch, pkg)
	if err != nil {
		log.Fatalln(err)
	}
}
