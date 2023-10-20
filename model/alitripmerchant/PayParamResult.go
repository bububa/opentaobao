package alitripmerchant

import (
	"sync"
)

// PayParamResult 结构体
type PayParamResult struct {
	// 签名
	PaySign string `json:"pay_sign,omitempty" xml:"pay_sign,omitempty"`
	// 签名方式
	SignType string `json:"sign_type,omitempty" xml:"sign_type,omitempty"`
	// 数据包
	PackageStr string `json:"package_str,omitempty" xml:"package_str,omitempty"`
	// 随机字符串
	NonceStr string `json:"nonce_str,omitempty" xml:"nonce_str,omitempty"`
	// 时间戳
	TimeStamp string `json:"time_stamp,omitempty" xml:"time_stamp,omitempty"`
	// 小程序id
	AppId string `json:"app_id,omitempty" xml:"app_id,omitempty"`
	// 订单编号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 预订结果
	BookResult bool `json:"book_result,omitempty" xml:"book_result,omitempty"`
}

var poolPayParamResult = sync.Pool{
	New: func() any {
		return new(PayParamResult)
	},
}

// GetPayParamResult() 从对象池中获取PayParamResult
func GetPayParamResult() *PayParamResult {
	return poolPayParamResult.Get().(*PayParamResult)
}

// ReleasePayParamResult 释放PayParamResult
func ReleasePayParamResult(v *PayParamResult) {
	v.PaySign = ""
	v.SignType = ""
	v.PackageStr = ""
	v.NonceStr = ""
	v.TimeStamp = ""
	v.AppId = ""
	v.OrderCode = ""
	v.BookResult = false
	poolPayParamResult.Put(v)
}
