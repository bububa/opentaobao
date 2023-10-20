package alihealthoutflow

import (
	"sync"
)

// PayTaxNoticeRequest 结构体
type PayTaxNoticeRequest struct {
	// 审核操作类型
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// appKey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 来源
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 扩展信息
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 医生id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

var poolPayTaxNoticeRequest = sync.Pool{
	New: func() any {
		return new(PayTaxNoticeRequest)
	},
}

// GetPayTaxNoticeRequest() 从对象池中获取PayTaxNoticeRequest
func GetPayTaxNoticeRequest() *PayTaxNoticeRequest {
	return poolPayTaxNoticeRequest.Get().(*PayTaxNoticeRequest)
}

// ReleasePayTaxNoticeRequest 释放PayTaxNoticeRequest
func ReleasePayTaxNoticeRequest(v *PayTaxNoticeRequest) {
	v.Action = ""
	v.AppKey = ""
	v.Source = ""
	v.ExtInfo = ""
	v.UserId = 0
	poolPayTaxNoticeRequest.Put(v)
}
