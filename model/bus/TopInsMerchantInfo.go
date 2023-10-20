package bus

import (
	"sync"
)

// TopInsMerchantInfo 结构体
type TopInsMerchantInfo struct {
	// 商户ID
	MerchantId string `json:"merchant_id,omitempty" xml:"merchant_id,omitempty"`
	// 商户名称
	MerchantName string `json:"merchant_name,omitempty" xml:"merchant_name,omitempty"`
}

var poolTopInsMerchantInfo = sync.Pool{
	New: func() any {
		return new(TopInsMerchantInfo)
	},
}

// GetTopInsMerchantInfo() 从对象池中获取TopInsMerchantInfo
func GetTopInsMerchantInfo() *TopInsMerchantInfo {
	return poolTopInsMerchantInfo.Get().(*TopInsMerchantInfo)
}

// ReleaseTopInsMerchantInfo 释放TopInsMerchantInfo
func ReleaseTopInsMerchantInfo(v *TopInsMerchantInfo) {
	v.MerchantId = ""
	v.MerchantName = ""
	poolTopInsMerchantInfo.Put(v)
}
