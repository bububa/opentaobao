package main

import "github.com/bububa/opentaobao/metadata"

type ApiPkg struct {
	metadata.PkgConfig
	Link string
}

type ApiPkgSlice []ApiPkg

func (p ApiPkgSlice) Len() int           { return len(p) }
func (p ApiPkgSlice) Less(i, j int) bool { return p[i].Id < p[j].Id }
func (p ApiPkgSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
