package ascpchannel

import (
	"sync"
)

// AlibabaAscpIndustryWaybillPreAcceptData 结构体
type AlibabaAscpIndustryWaybillPreAcceptData struct {
	// 运单号
	ExpressNo string `json:"express_no,omitempty" xml:"express_no,omitempty"`
	// 物流品牌
	LogisticsBrand string `json:"logistics_brand,omitempty" xml:"logistics_brand,omitempty"`
}

var poolAlibabaAscpIndustryWaybillPreAcceptData = sync.Pool{
	New: func() any {
		return new(AlibabaAscpIndustryWaybillPreAcceptData)
	},
}

// GetAlibabaAscpIndustryWaybillPreAcceptData() 从对象池中获取AlibabaAscpIndustryWaybillPreAcceptData
func GetAlibabaAscpIndustryWaybillPreAcceptData() *AlibabaAscpIndustryWaybillPreAcceptData {
	return poolAlibabaAscpIndustryWaybillPreAcceptData.Get().(*AlibabaAscpIndustryWaybillPreAcceptData)
}

// ReleaseAlibabaAscpIndustryWaybillPreAcceptData 释放AlibabaAscpIndustryWaybillPreAcceptData
func ReleaseAlibabaAscpIndustryWaybillPreAcceptData(v *AlibabaAscpIndustryWaybillPreAcceptData) {
	v.ExpressNo = ""
	v.LogisticsBrand = ""
	poolAlibabaAscpIndustryWaybillPreAcceptData.Put(v)
}
