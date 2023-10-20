package logistic

import (
	"sync"
)

// TmsCollectResponse 结构体
type TmsCollectResponse struct {
	// 风控通过重量
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
	// 重量是否通过风控
	WeightCheckResult bool `json:"weight_check_result,omitempty" xml:"weight_check_result,omitempty"`
}

var poolTmsCollectResponse = sync.Pool{
	New: func() any {
		return new(TmsCollectResponse)
	},
}

// GetTmsCollectResponse() 从对象池中获取TmsCollectResponse
func GetTmsCollectResponse() *TmsCollectResponse {
	return poolTmsCollectResponse.Get().(*TmsCollectResponse)
}

// ReleaseTmsCollectResponse 释放TmsCollectResponse
func ReleaseTmsCollectResponse(v *TmsCollectResponse) {
	v.Weight = ""
	v.WeightCheckResult = false
	poolTmsCollectResponse.Put(v)
}
