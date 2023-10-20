package omniorder

import (
	"sync"
)

// BaseResult 结构体
type BaseResult struct {
	// 返回的数据实体
	CommissionResultList []CommissionResultDto `json:"commission_result_list,omitempty" xml:"commission_result_list>commission_result_dto,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回的执行状态吗
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 总记录数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 是否执行成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBaseResult = sync.Pool{
	New: func() any {
		return new(BaseResult)
	},
}

// GetBaseResult() 从对象池中获取BaseResult
func GetBaseResult() *BaseResult {
	return poolBaseResult.Get().(*BaseResult)
}

// ReleaseBaseResult 释放BaseResult
func ReleaseBaseResult(v *BaseResult) {
	v.CommissionResultList = v.CommissionResultList[:0]
	v.Message = ""
	v.Code = ""
	v.Total = 0
	v.PageSize = 0
	v.PageNo = 0
	v.Success = false
	poolBaseResult.Put(v)
}
