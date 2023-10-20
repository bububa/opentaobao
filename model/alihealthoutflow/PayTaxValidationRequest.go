package alihealthoutflow

import (
	"sync"
)

// PayTaxValidationRequest 结构体
type PayTaxValidationRequest struct {
	// 信息获取凭证
	Ticket string `json:"ticket,omitempty" xml:"ticket,omitempty"`
	// appKey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 医生id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

var poolPayTaxValidationRequest = sync.Pool{
	New: func() any {
		return new(PayTaxValidationRequest)
	},
}

// GetPayTaxValidationRequest() 从对象池中获取PayTaxValidationRequest
func GetPayTaxValidationRequest() *PayTaxValidationRequest {
	return poolPayTaxValidationRequest.Get().(*PayTaxValidationRequest)
}

// ReleasePayTaxValidationRequest 释放PayTaxValidationRequest
func ReleasePayTaxValidationRequest(v *PayTaxValidationRequest) {
	v.Ticket = ""
	v.AppKey = ""
	v.UserId = 0
	poolPayTaxValidationRequest.Put(v)
}
