package cainiaoncwl

import (
	"sync"
)

// JhReceiverInfo 结构体
type JhReceiverInfo struct {
	// 收货城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 收货县或者区
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 收货省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 收货明细地址
	ReceiverAddress string `json:"receiver_address,omitempty" xml:"receiver_address,omitempty"`
	// 收货手机号
	ReceiverMobile string `json:"receiver_mobile,omitempty" xml:"receiver_mobile,omitempty"`
	// 收货姓名
	ReceiverName string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
	// 收货电话
	ReceiverPhone string `json:"receiver_phone,omitempty" xml:"receiver_phone,omitempty"`
	// 收货邮编
	ReceiverZip string `json:"receiver_zip,omitempty" xml:"receiver_zip,omitempty"`
	// 收货镇或者街区
	Street string `json:"street,omitempty" xml:"street,omitempty"`
}

var poolJhReceiverInfo = sync.Pool{
	New: func() any {
		return new(JhReceiverInfo)
	},
}

// GetJhReceiverInfo() 从对象池中获取JhReceiverInfo
func GetJhReceiverInfo() *JhReceiverInfo {
	return poolJhReceiverInfo.Get().(*JhReceiverInfo)
}

// ReleaseJhReceiverInfo 释放JhReceiverInfo
func ReleaseJhReceiverInfo(v *JhReceiverInfo) {
	v.City = ""
	v.District = ""
	v.Province = ""
	v.ReceiverAddress = ""
	v.ReceiverMobile = ""
	v.ReceiverName = ""
	v.ReceiverPhone = ""
	v.ReceiverZip = ""
	v.Street = ""
	poolJhReceiverInfo.Put(v)
}
