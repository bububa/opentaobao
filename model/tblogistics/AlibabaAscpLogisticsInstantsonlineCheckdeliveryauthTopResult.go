package tblogistics

import (
	"sync"
)

// AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthTopResult 结构体
type AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthTopResult struct {
	// 数据
	Data *CheckDeliveryAuthTopResponse `json:"data,omitempty" xml:"data,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAscpLogisticsInstantsonlineCheckdeliveryauthTopResult = sync.Pool{
	New: func() any {
		return new(AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthTopResult)
	},
}

// GetAlibabaAscpLogisticsInstantsonlineCheckdeliveryauthTopResult() 从对象池中获取AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthTopResult
func GetAlibabaAscpLogisticsInstantsonlineCheckdeliveryauthTopResult() *AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthTopResult {
	return poolAlibabaAscpLogisticsInstantsonlineCheckdeliveryauthTopResult.Get().(*AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthTopResult)
}

// ReleaseAlibabaAscpLogisticsInstantsonlineCheckdeliveryauthTopResult 释放AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthTopResult
func ReleaseAlibabaAscpLogisticsInstantsonlineCheckdeliveryauthTopResult(v *AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthTopResult) {
	v.Data = nil
	v.Success = false
	poolAlibabaAscpLogisticsInstantsonlineCheckdeliveryauthTopResult.Put(v)
}
