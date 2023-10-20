package ascpchannel

import (
	"sync"
)

// AlibabaAscpIndustryUopSupplierConsignoderData 结构体
type AlibabaAscpIndustryUopSupplierConsignoderData struct {
	// 送装物流单号
	ExpressCode string `json:"express_code,omitempty" xml:"express_code,omitempty"`
}

var poolAlibabaAscpIndustryUopSupplierConsignoderData = sync.Pool{
	New: func() any {
		return new(AlibabaAscpIndustryUopSupplierConsignoderData)
	},
}

// GetAlibabaAscpIndustryUopSupplierConsignoderData() 从对象池中获取AlibabaAscpIndustryUopSupplierConsignoderData
func GetAlibabaAscpIndustryUopSupplierConsignoderData() *AlibabaAscpIndustryUopSupplierConsignoderData {
	return poolAlibabaAscpIndustryUopSupplierConsignoderData.Get().(*AlibabaAscpIndustryUopSupplierConsignoderData)
}

// ReleaseAlibabaAscpIndustryUopSupplierConsignoderData 释放AlibabaAscpIndustryUopSupplierConsignoderData
func ReleaseAlibabaAscpIndustryUopSupplierConsignoderData(v *AlibabaAscpIndustryUopSupplierConsignoderData) {
	v.ExpressCode = ""
	poolAlibabaAscpIndustryUopSupplierConsignoderData.Put(v)
}
