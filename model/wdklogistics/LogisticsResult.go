package wdklogistics

import (
	"sync"
)

// LogisticsResult 结构体
type LogisticsResult struct {
	// 内容节点
	Datas []AlibabaWdkLogisticsPusPickupCararrivedData `json:"datas,omitempty" xml:"datas>alibaba_wdk_logistics_pus_pickup_cararrived_data,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 接口服务查询结果描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolLogisticsResult = sync.Pool{
	New: func() any {
		return new(LogisticsResult)
	},
}

// GetLogisticsResult() 从对象池中获取LogisticsResult
func GetLogisticsResult() *LogisticsResult {
	return poolLogisticsResult.Get().(*LogisticsResult)
}

// ReleaseLogisticsResult 释放LogisticsResult
func ReleaseLogisticsResult(v *LogisticsResult) {
	v.Datas = v.Datas[:0]
	v.Code = ""
	v.Message = ""
	v.Success = false
	poolLogisticsResult.Put(v)
}
