package inventory

import (
	"sync"
)

// SingleResult 结构体
type SingleResult struct {
	// data
	AdjustResults []InventoryCheckResultDto `json:"adjust_results,omitempty" xml:"adjust_results>inventory_check_result_dto,omitempty"`
	// 地点关系
	LocationRelationList []LocationRelationDto `json:"location_relation_list,omitempty" xml:"location_relation_list>location_relation_dto,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 如果是失败，可能是部分失败。如果是成功，则全部成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolSingleResult = sync.Pool{
	New: func() any {
		return new(SingleResult)
	},
}

// GetSingleResult() 从对象池中获取SingleResult
func GetSingleResult() *SingleResult {
	return poolSingleResult.Get().(*SingleResult)
}

// ReleaseSingleResult 释放SingleResult
func ReleaseSingleResult(v *SingleResult) {
	v.AdjustResults = v.AdjustResults[:0]
	v.LocationRelationList = v.LocationRelationList[:0]
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Success = false
	poolSingleResult.Put(v)
}
