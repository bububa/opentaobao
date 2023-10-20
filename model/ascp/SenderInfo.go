package ascp

import (
	"sync"
)

// SenderInfo 结构体
type SenderInfo struct {
	// 公司名称
	Company string `json:"company,omitempty" xml:"company,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 邮编
	ZipCode string `json:"zip_code,omitempty" xml:"zip_code,omitempty"`
	// 固定电话
	Tel string `json:"tel,omitempty" xml:"tel,omitempty"`
	// 移动电话
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 电子邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 国家二字码
	CountryCode string `json:"country_code,omitempty" xml:"country_code,omitempty"`
	// 省份
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 区域
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 村镇
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 详细地址
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
}

var poolSenderInfo = sync.Pool{
	New: func() any {
		return new(SenderInfo)
	},
}

// GetSenderInfo() 从对象池中获取SenderInfo
func GetSenderInfo() *SenderInfo {
	return poolSenderInfo.Get().(*SenderInfo)
}

// ReleaseSenderInfo 释放SenderInfo
func ReleaseSenderInfo(v *SenderInfo) {
	v.Company = ""
	v.Name = ""
	v.ZipCode = ""
	v.Tel = ""
	v.Mobile = ""
	v.Email = ""
	v.CountryCode = ""
	v.Province = ""
	v.City = ""
	v.Area = ""
	v.Town = ""
	v.DetailAddress = ""
	poolSenderInfo.Put(v)
}
