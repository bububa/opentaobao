package alitripmerchant

import (
	"sync"
)

// OrderPayDto 结构体
type OrderPayDto struct {
	// 签名
	PaySign string `json:"pay_sign,omitempty" xml:"pay_sign,omitempty"`
	// 签名类型
	SignType string `json:"sign_type,omitempty" xml:"sign_type,omitempty"`
	// 数据包
	PackageStr string `json:"package_str,omitempty" xml:"package_str,omitempty"`
	// 随机字符串
	NonceStr string `json:"nonce_str,omitempty" xml:"nonce_str,omitempty"`
	// 时间戳
	TimeStamp string `json:"time_stamp,omitempty" xml:"time_stamp,omitempty"`
	// 小程序id
	AppId string `json:"app_id,omitempty" xml:"app_id,omitempty"`
	// 预订结果
	BookResult bool `json:"book_result,omitempty" xml:"book_result,omitempty"`
}

var poolOrderPayDto = sync.Pool{
	New: func() any {
		return new(OrderPayDto)
	},
}

// GetOrderPayDto() 从对象池中获取OrderPayDto
func GetOrderPayDto() *OrderPayDto {
	return poolOrderPayDto.Get().(*OrderPayDto)
}

// ReleaseOrderPayDto 释放OrderPayDto
func ReleaseOrderPayDto(v *OrderPayDto) {
	v.PaySign = ""
	v.SignType = ""
	v.PackageStr = ""
	v.NonceStr = ""
	v.TimeStamp = ""
	v.AppId = ""
	v.BookResult = false
	poolOrderPayDto.Put(v)
}
