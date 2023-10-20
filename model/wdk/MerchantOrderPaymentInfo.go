package wdk

import (
	"sync"
)

// MerchantOrderPaymentInfo 结构体
type MerchantOrderPaymentInfo struct {
	// 其他支付方式支付后获得userId，例如：微信平台的openId
	Ouid string `json:"ouid,omitempty" xml:"ouid,omitempty"`
	// 线上支付订单号，线下流水号，代金券/优惠券为优惠券实例id
	SerialNum string `json:"serial_num,omitempty" xml:"serial_num,omitempty"`
	// 支付宝支付方式后获得的userId
	Tuid string `json:"tuid,omitempty" xml:"tuid,omitempty"`
	// 支付类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 付款金额
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

var poolMerchantOrderPaymentInfo = sync.Pool{
	New: func() any {
		return new(MerchantOrderPaymentInfo)
	},
}

// GetMerchantOrderPaymentInfo() 从对象池中获取MerchantOrderPaymentInfo
func GetMerchantOrderPaymentInfo() *MerchantOrderPaymentInfo {
	return poolMerchantOrderPaymentInfo.Get().(*MerchantOrderPaymentInfo)
}

// ReleaseMerchantOrderPaymentInfo 释放MerchantOrderPaymentInfo
func ReleaseMerchantOrderPaymentInfo(v *MerchantOrderPaymentInfo) {
	v.Ouid = ""
	v.SerialNum = ""
	v.Tuid = ""
	v.Type = ""
	v.Amount = 0
	poolMerchantOrderPaymentInfo.Put(v)
}
