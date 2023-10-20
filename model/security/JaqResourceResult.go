package security

import (
	"sync"
)

// JaqResourceResult 结构体
type JaqResourceResult struct {
	// 请求事件唯一标识
	EventId string `json:"event_id,omitempty" xml:"event_id,omitempty"`
	// 资源包的md5
	Md5 string `json:"md5,omitempty" xml:"md5,omitempty"`
	// 资源的cdn下载链接
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 资源版本号
	Version string `json:"version,omitempty" xml:"version,omitempty"`
}

var poolJaqResourceResult = sync.Pool{
	New: func() any {
		return new(JaqResourceResult)
	},
}

// GetJaqResourceResult() 从对象池中获取JaqResourceResult
func GetJaqResourceResult() *JaqResourceResult {
	return poolJaqResourceResult.Get().(*JaqResourceResult)
}

// ReleaseJaqResourceResult 释放JaqResourceResult
func ReleaseJaqResourceResult(v *JaqResourceResult) {
	v.EventId = ""
	v.Md5 = ""
	v.Url = ""
	v.Version = ""
	poolJaqResourceResult.Put(v)
}
