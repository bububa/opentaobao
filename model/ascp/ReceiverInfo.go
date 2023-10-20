package ascp

import (
	"sync"
)

// ReceiverInfo 结构体
type ReceiverInfo struct {
	// 收件人国家(地区)二字码。https://www.ufsoo.com/news/detail-59307681-b5e1-4328-a6e6-20fddd6c5ec6.html
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 收件人所在省份
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 收件人所在城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 收件人所在地区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 收件人所在村镇/街道
	Town string `json:"town,omitempty" xml:"town,omitempty"`
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
	// 收件人证件类型(1-身份证、2-军官证、3-护照、4-其他)
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 国家二字码
	CountryCode string `json:"country_code,omitempty" xml:"country_code,omitempty"`
	// 详细地址
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 收件人地址ID
	Oaid string `json:"oaid,omitempty" xml:"oaid,omitempty"`
	// 是否虚拟号
	Privacy string `json:"privacy,omitempty" xml:"privacy,omitempty"`
}

var poolReceiverInfo = sync.Pool{
	New: func() any {
		return new(ReceiverInfo)
	},
}

// GetReceiverInfo() 从对象池中获取ReceiverInfo
func GetReceiverInfo() *ReceiverInfo {
	return poolReceiverInfo.Get().(*ReceiverInfo)
}

// ReleaseReceiverInfo 释放ReceiverInfo
func ReleaseReceiverInfo(v *ReceiverInfo) {
	v.Country = ""
	v.Province = ""
	v.City = ""
	v.Area = ""
	v.Town = ""
	v.Company = ""
	v.Name = ""
	v.ZipCode = ""
	v.Tel = ""
	v.Mobile = ""
	v.Email = ""
	v.CountryCode = ""
	v.DetailAddress = ""
	v.Oaid = ""
	v.Privacy = ""
	poolReceiverInfo.Put(v)
}
