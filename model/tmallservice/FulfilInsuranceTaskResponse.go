package tmallservice

import (
	"sync"
)

// FulfilInsuranceTaskResponse 结构体
type FulfilInsuranceTaskResponse struct {
	// 外部单号id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 履行单id
	FulfilTaskId int64 `json:"fulfil_task_id,omitempty" xml:"fulfil_task_id,omitempty"`
	// 工单id
	WorkcardId int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
}

var poolFulfilInsuranceTaskResponse = sync.Pool{
	New: func() any {
		return new(FulfilInsuranceTaskResponse)
	},
}

// GetFulfilInsuranceTaskResponse() 从对象池中获取FulfilInsuranceTaskResponse
func GetFulfilInsuranceTaskResponse() *FulfilInsuranceTaskResponse {
	return poolFulfilInsuranceTaskResponse.Get().(*FulfilInsuranceTaskResponse)
}

// ReleaseFulfilInsuranceTaskResponse 释放FulfilInsuranceTaskResponse
func ReleaseFulfilInsuranceTaskResponse(v *FulfilInsuranceTaskResponse) {
	v.OuterId = ""
	v.FulfilTaskId = 0
	v.WorkcardId = 0
	poolFulfilInsuranceTaskResponse.Put(v)
}
