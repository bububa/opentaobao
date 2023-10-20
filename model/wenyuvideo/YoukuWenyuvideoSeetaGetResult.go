package wenyuvideo

import (
	"sync"
)

// YoukuWenyuvideoSeetaGetResult 结构体
type YoukuWenyuvideoSeetaGetResult struct {
	// 返回数据
	Values []YoukuWenyuvideoSeetaGetModel `json:"values,omitempty" xml:"values>youku_wenyuvideo_seeta_get_model,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 结果代码
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolYoukuWenyuvideoSeetaGetResult = sync.Pool{
	New: func() any {
		return new(YoukuWenyuvideoSeetaGetResult)
	},
}

// GetYoukuWenyuvideoSeetaGetResult() 从对象池中获取YoukuWenyuvideoSeetaGetResult
func GetYoukuWenyuvideoSeetaGetResult() *YoukuWenyuvideoSeetaGetResult {
	return poolYoukuWenyuvideoSeetaGetResult.Get().(*YoukuWenyuvideoSeetaGetResult)
}

// ReleaseYoukuWenyuvideoSeetaGetResult 释放YoukuWenyuvideoSeetaGetResult
func ReleaseYoukuWenyuvideoSeetaGetResult(v *YoukuWenyuvideoSeetaGetResult) {
	v.Values = v.Values[:0]
	v.MsgCode = ""
	v.MsgInfo = ""
	v.HttpStatusCode = 0
	v.Success = false
	poolYoukuWenyuvideoSeetaGetResult.Put(v)
}
