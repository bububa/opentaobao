package alsc

import (
	"sync"
)

// QueryCodeBizAttrRequest 结构体
type QueryCodeBizAttrRequest struct {
	// 码使用的业务场景
	BizScenePrefix string `json:"biz_scene_prefix,omitempty" xml:"biz_scene_prefix,omitempty"`
	// 码值
	CodeValue string `json:"code_value,omitempty" xml:"code_value,omitempty"`
}

var poolQueryCodeBizAttrRequest = sync.Pool{
	New: func() any {
		return new(QueryCodeBizAttrRequest)
	},
}

// GetQueryCodeBizAttrRequest() 从对象池中获取QueryCodeBizAttrRequest
func GetQueryCodeBizAttrRequest() *QueryCodeBizAttrRequest {
	return poolQueryCodeBizAttrRequest.Get().(*QueryCodeBizAttrRequest)
}

// ReleaseQueryCodeBizAttrRequest 释放QueryCodeBizAttrRequest
func ReleaseQueryCodeBizAttrRequest(v *QueryCodeBizAttrRequest) {
	v.BizScenePrefix = ""
	v.CodeValue = ""
	poolQueryCodeBizAttrRequest.Put(v)
}
