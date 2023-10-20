package alsc

import (
	"sync"
)

// ResetPayPasswdOpenReq 结构体
type ResetPayPasswdOpenReq struct {
	// 品牌ID / 外部品牌id  2选1
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 操作人ID
	OperatorId string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	// 操作人名称
	OperatorName string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
	// 请求ID，幂等处理
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 20180988753656
	CustomerId string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
	// 外部品牌id
	OutBrandId string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
}

var poolResetPayPasswdOpenReq = sync.Pool{
	New: func() any {
		return new(ResetPayPasswdOpenReq)
	},
}

// GetResetPayPasswdOpenReq() 从对象池中获取ResetPayPasswdOpenReq
func GetResetPayPasswdOpenReq() *ResetPayPasswdOpenReq {
	return poolResetPayPasswdOpenReq.Get().(*ResetPayPasswdOpenReq)
}

// ReleaseResetPayPasswdOpenReq 释放ResetPayPasswdOpenReq
func ReleaseResetPayPasswdOpenReq(v *ResetPayPasswdOpenReq) {
	v.BrandId = ""
	v.OperatorId = ""
	v.OperatorName = ""
	v.RequestId = ""
	v.CustomerId = ""
	v.OutBrandId = ""
	poolResetPayPasswdOpenReq.Put(v)
}
