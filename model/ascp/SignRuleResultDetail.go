package ascp

import (
	"sync"
)

// SignRuleResultDetail 结构体
type SignRuleResultDetail struct {
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 收货地：省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 收货地：城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 收货地：区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 收货地：街道
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// wms货主id
	WmsOwnerCode string `json:"wms_owner_code,omitempty" xml:"wms_owner_code,omitempty"`
}

var poolSignRuleResultDetail = sync.Pool{
	New: func() any {
		return new(SignRuleResultDetail)
	},
}

// GetSignRuleResultDetail() 从对象池中获取SignRuleResultDetail
func GetSignRuleResultDetail() *SignRuleResultDetail {
	return poolSignRuleResultDetail.Get().(*SignRuleResultDetail)
}

// ReleaseSignRuleResultDetail 释放SignRuleResultDetail
func ReleaseSignRuleResultDetail(v *SignRuleResultDetail) {
	v.Code = ""
	v.Message = ""
	v.Province = ""
	v.City = ""
	v.Area = ""
	v.Town = ""
	v.WmsOwnerCode = ""
	poolSignRuleResultDetail.Put(v)
}
