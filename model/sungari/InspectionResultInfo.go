package sungari

import (
	"sync"
)

// InspectionResultInfo 结构体
type InspectionResultInfo struct {
	// 卖家nick
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 发货地址
	SendAddress string `json:"send_address,omitempty" xml:"send_address,omitempty"`
	// 注册地址
	RegisterAddress string `json:"register_address,omitempty" xml:"register_address,omitempty"`
	// 营业执照编号
	LicenceNo string `json:"licence_no,omitempty" xml:"licence_no,omitempty"`
	// 卖家电话
	SellerTel string `json:"seller_tel,omitempty" xml:"seller_tel,omitempty"`
	// 认证名称
	CertificationName string `json:"certification_name,omitempty" xml:"certification_name,omitempty"`
	// 处置结果
	DisposeResult string `json:"dispose_result,omitempty" xml:"dispose_result,omitempty"`
}

var poolInspectionResultInfo = sync.Pool{
	New: func() any {
		return new(InspectionResultInfo)
	},
}

// GetInspectionResultInfo() 从对象池中获取InspectionResultInfo
func GetInspectionResultInfo() *InspectionResultInfo {
	return poolInspectionResultInfo.Get().(*InspectionResultInfo)
}

// ReleaseInspectionResultInfo 释放InspectionResultInfo
func ReleaseInspectionResultInfo(v *InspectionResultInfo) {
	v.SellerNick = ""
	v.SendAddress = ""
	v.RegisterAddress = ""
	v.LicenceNo = ""
	v.SellerTel = ""
	v.CertificationName = ""
	v.DisposeResult = ""
	poolInspectionResultInfo.Put(v)
}
