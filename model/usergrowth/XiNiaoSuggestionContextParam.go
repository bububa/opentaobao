package usergrowth

import (
	"sync"
)

// XiNiaoSuggestionContextParam 结构体
type XiNiaoSuggestionContextParam struct {
	// open_id
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// open_id
	OpenId2 string `json:"open_id2,omitempty" xml:"open_id2,omitempty"`
	// 站点id
	UnionCode string `json:"union_code,omitempty" xml:"union_code,omitempty"`
	// 包裹数量
	PackageNum int64 `json:"package_num,omitempty" xml:"package_num,omitempty"`
}

var poolXiNiaoSuggestionContextParam = sync.Pool{
	New: func() any {
		return new(XiNiaoSuggestionContextParam)
	},
}

// GetXiNiaoSuggestionContextParam() 从对象池中获取XiNiaoSuggestionContextParam
func GetXiNiaoSuggestionContextParam() *XiNiaoSuggestionContextParam {
	return poolXiNiaoSuggestionContextParam.Get().(*XiNiaoSuggestionContextParam)
}

// ReleaseXiNiaoSuggestionContextParam 释放XiNiaoSuggestionContextParam
func ReleaseXiNiaoSuggestionContextParam(v *XiNiaoSuggestionContextParam) {
	v.OpenId = ""
	v.OpenId2 = ""
	v.UnionCode = ""
	v.PackageNum = 0
	poolXiNiaoSuggestionContextParam.Put(v)
}
