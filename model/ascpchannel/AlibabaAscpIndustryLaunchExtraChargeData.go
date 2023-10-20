package ascpchannel

import (
	"sync"
)

// AlibabaAscpIndustryLaunchExtraChargeData 结构体
type AlibabaAscpIndustryLaunchExtraChargeData struct {
	// 增加费用服务调整单ID
	ExtraChargeServiceOrderId string `json:"extra_charge_service_order_id,omitempty" xml:"extra_charge_service_order_id,omitempty"`
	// 扩展字段
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
}

var poolAlibabaAscpIndustryLaunchExtraChargeData = sync.Pool{
	New: func() any {
		return new(AlibabaAscpIndustryLaunchExtraChargeData)
	},
}

// GetAlibabaAscpIndustryLaunchExtraChargeData() 从对象池中获取AlibabaAscpIndustryLaunchExtraChargeData
func GetAlibabaAscpIndustryLaunchExtraChargeData() *AlibabaAscpIndustryLaunchExtraChargeData {
	return poolAlibabaAscpIndustryLaunchExtraChargeData.Get().(*AlibabaAscpIndustryLaunchExtraChargeData)
}

// ReleaseAlibabaAscpIndustryLaunchExtraChargeData 释放AlibabaAscpIndustryLaunchExtraChargeData
func ReleaseAlibabaAscpIndustryLaunchExtraChargeData(v *AlibabaAscpIndustryLaunchExtraChargeData) {
	v.ExtraChargeServiceOrderId = ""
	v.Feature = ""
	poolAlibabaAscpIndustryLaunchExtraChargeData.Put(v)
}
