package logistic

import (
	"sync"
)

// AlibabaEleFengniaoChainstoreContractCancelData 结构体
type AlibabaEleFengniaoChainstoreContractCancelData struct {
	// 门店code
	ChainstoreCodes []string `json:"chainstore_codes,omitempty" xml:"chainstore_codes>string,omitempty"`
	// appid
	AppId string `json:"app_id,omitempty" xml:"app_id,omitempty"`
	// 商户code
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
}

var poolAlibabaEleFengniaoChainstoreContractCancelData = sync.Pool{
	New: func() any {
		return new(AlibabaEleFengniaoChainstoreContractCancelData)
	},
}

// GetAlibabaEleFengniaoChainstoreContractCancelData() 从对象池中获取AlibabaEleFengniaoChainstoreContractCancelData
func GetAlibabaEleFengniaoChainstoreContractCancelData() *AlibabaEleFengniaoChainstoreContractCancelData {
	return poolAlibabaEleFengniaoChainstoreContractCancelData.Get().(*AlibabaEleFengniaoChainstoreContractCancelData)
}

// ReleaseAlibabaEleFengniaoChainstoreContractCancelData 释放AlibabaEleFengniaoChainstoreContractCancelData
func ReleaseAlibabaEleFengniaoChainstoreContractCancelData(v *AlibabaEleFengniaoChainstoreContractCancelData) {
	v.ChainstoreCodes = v.ChainstoreCodes[:0]
	v.AppId = ""
	v.MerchantCode = ""
	poolAlibabaEleFengniaoChainstoreContractCancelData.Put(v)
}
