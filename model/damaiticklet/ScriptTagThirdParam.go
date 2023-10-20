package damaiticklet

import (
	"sync"
)

// ScriptTagThirdParam 结构体
type ScriptTagThirdParam struct {
	// 剧本名称
	TagName string `json:"tag_name,omitempty" xml:"tag_name,omitempty"`
	// 剧本类型,核心题材1   剧本题材 2  时代背景 3   流派/类型 4  剧本类型 5
	TagType int64 `json:"tag_type,omitempty" xml:"tag_type,omitempty"`
	// 剧本标签id
	OutTagId int64 `json:"out_tag_id,omitempty" xml:"out_tag_id,omitempty"`
}

var poolScriptTagThirdParam = sync.Pool{
	New: func() any {
		return new(ScriptTagThirdParam)
	},
}

// GetScriptTagThirdParam() 从对象池中获取ScriptTagThirdParam
func GetScriptTagThirdParam() *ScriptTagThirdParam {
	return poolScriptTagThirdParam.Get().(*ScriptTagThirdParam)
}

// ReleaseScriptTagThirdParam 释放ScriptTagThirdParam
func ReleaseScriptTagThirdParam(v *ScriptTagThirdParam) {
	v.TagName = ""
	v.TagType = 0
	v.OutTagId = 0
	poolScriptTagThirdParam.Put(v)
}
