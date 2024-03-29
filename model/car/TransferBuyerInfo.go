package car

import (
	"sync"
)

// TransferBuyerInfo 结构体
type TransferBuyerInfo struct {
	// 飞猪虚拟号
	TravellerSecretPhone string `json:"traveller_secret_phone,omitempty" xml:"traveller_secret_phone,omitempty"`
	// 乘客真实号(后四位)
	PassengerRealPhoneLast string `json:"passenger_real_phone_last,omitempty" xml:"passenger_real_phone_last,omitempty"`
	// 虚拟号失效时间
	SecretPhoneEndTime string `json:"secret_phone_end_time,omitempty" xml:"secret_phone_end_time,omitempty"`
	// 乘客真实号（可能为空）
	PassengerRealPhone string `json:"passenger_real_phone,omitempty" xml:"passenger_real_phone,omitempty"`
}

var poolTransferBuyerInfo = sync.Pool{
	New: func() any {
		return new(TransferBuyerInfo)
	},
}

// GetTransferBuyerInfo() 从对象池中获取TransferBuyerInfo
func GetTransferBuyerInfo() *TransferBuyerInfo {
	return poolTransferBuyerInfo.Get().(*TransferBuyerInfo)
}

// ReleaseTransferBuyerInfo 释放TransferBuyerInfo
func ReleaseTransferBuyerInfo(v *TransferBuyerInfo) {
	v.TravellerSecretPhone = ""
	v.PassengerRealPhoneLast = ""
	v.SecretPhoneEndTime = ""
	v.PassengerRealPhone = ""
	poolTransferBuyerInfo.Put(v)
}
