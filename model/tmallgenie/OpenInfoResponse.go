package tmallgenie

import (
	"sync"
)

// OpenInfoResponse 结构体
type OpenInfoResponse struct {
	// 关联id
	UnionIds []UnionIdInfo `json:"union_ids,omitempty" xml:"union_ids>union_id_info,omitempty"`
	// 开放openId
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 入参内容
	Param *ConverterIdRequest `json:"param,omitempty" xml:"param,omitempty"`
}

var poolOpenInfoResponse = sync.Pool{
	New: func() any {
		return new(OpenInfoResponse)
	},
}

// GetOpenInfoResponse() 从对象池中获取OpenInfoResponse
func GetOpenInfoResponse() *OpenInfoResponse {
	return poolOpenInfoResponse.Get().(*OpenInfoResponse)
}

// ReleaseOpenInfoResponse 释放OpenInfoResponse
func ReleaseOpenInfoResponse(v *OpenInfoResponse) {
	v.UnionIds = v.UnionIds[:0]
	v.OpenId = ""
	v.Param = nil
	poolOpenInfoResponse.Put(v)
}
