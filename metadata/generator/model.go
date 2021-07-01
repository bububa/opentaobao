package main

import "github.com/bububa/opentaobao/metadata"

// SDK 包配置结构体，用于生成README中包列表
type ApiPkg struct {
	metadata.PkgConfig
	Link string
}

// SDK 包配置结构体按类目ID排序
type ApiPkgSlice []ApiPkg

func (p ApiPkgSlice) Len() int           { return len(p) }
func (p ApiPkgSlice) Less(i, j int) bool { return p[i].Id < p[j].Id }
func (p ApiPkgSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
