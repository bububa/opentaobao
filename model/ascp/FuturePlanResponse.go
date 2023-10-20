package ascp

import (
	"sync"
)

// FuturePlanResponse 结构体
type FuturePlanResponse struct {
	// 负卖设置结果
	FuturePlanItemResultList []FuturePlanItemResult `json:"future_plan_item_result_list,omitempty" xml:"future_plan_item_result_list>future_plan_item_result,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 调用系统链路是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolFuturePlanResponse = sync.Pool{
	New: func() any {
		return new(FuturePlanResponse)
	},
}

// GetFuturePlanResponse() 从对象池中获取FuturePlanResponse
func GetFuturePlanResponse() *FuturePlanResponse {
	return poolFuturePlanResponse.Get().(*FuturePlanResponse)
}

// ReleaseFuturePlanResponse 释放FuturePlanResponse
func ReleaseFuturePlanResponse(v *FuturePlanResponse) {
	v.FuturePlanItemResultList = v.FuturePlanItemResultList[:0]
	v.ErrorCode = ""
	v.ErrorMessage = ""
	v.Success = false
	poolFuturePlanResponse.Put(v)
}
