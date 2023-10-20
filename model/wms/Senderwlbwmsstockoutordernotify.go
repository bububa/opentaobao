package wms

import (
	"sync"
)

// Senderwlbwmsstockoutordernotify 结构体
type Senderwlbwmsstockoutordernotify struct {
	// 发件方镇
	SenderTown string `json:"sender_town,omitempty" xml:"sender_town,omitempty"`
	// 发件方邮编
	SenderZipCode string `json:"sender_zip_code,omitempty" xml:"sender_zip_code,omitempty"`
	// 发件方地址
	SenderAddress string `json:"sender_address,omitempty" xml:"sender_address,omitempty"`
	// 发件方名称
	SenderName string `json:"sender_name,omitempty" xml:"sender_name,omitempty"`
	// 发件方省份
	SenderProvince string `json:"sender_province,omitempty" xml:"sender_province,omitempty"`
	// 发件方区县
	SenderArea string `json:"sender_area,omitempty" xml:"sender_area,omitempty"`
	// 发件方城市
	SenderCity string `json:"sender_city,omitempty" xml:"sender_city,omitempty"`
	// 发件方手机
	SenderMobile string `json:"sender_mobile,omitempty" xml:"sender_mobile,omitempty"`
	// 发件方电话
	SenderPhone string `json:"sender_phone,omitempty" xml:"sender_phone,omitempty"`
}

var poolSenderwlbwmsstockoutordernotify = sync.Pool{
	New: func() any {
		return new(Senderwlbwmsstockoutordernotify)
	},
}

// GetSenderwlbwmsstockoutordernotify() 从对象池中获取Senderwlbwmsstockoutordernotify
func GetSenderwlbwmsstockoutordernotify() *Senderwlbwmsstockoutordernotify {
	return poolSenderwlbwmsstockoutordernotify.Get().(*Senderwlbwmsstockoutordernotify)
}

// ReleaseSenderwlbwmsstockoutordernotify 释放Senderwlbwmsstockoutordernotify
func ReleaseSenderwlbwmsstockoutordernotify(v *Senderwlbwmsstockoutordernotify) {
	v.SenderTown = ""
	v.SenderZipCode = ""
	v.SenderAddress = ""
	v.SenderName = ""
	v.SenderProvince = ""
	v.SenderArea = ""
	v.SenderCity = ""
	v.SenderMobile = ""
	v.SenderPhone = ""
	poolSenderwlbwmsstockoutordernotify.Put(v)
}
