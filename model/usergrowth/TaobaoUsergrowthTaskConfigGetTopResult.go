package usergrowth

import (
	"sync"
)

// TaobaoUsergrowthTaskConfigGetTopResult 结构体
type TaobaoUsergrowthTaskConfigGetTopResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 任务配置
	Data *TaskConfig `json:"data,omitempty" xml:"data,omitempty"`
	// 是否执行成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoUsergrowthTaskConfigGetTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUsergrowthTaskConfigGetTopResult)
	},
}

// GetTaobaoUsergrowthTaskConfigGetTopResult() 从对象池中获取TaobaoUsergrowthTaskConfigGetTopResult
func GetTaobaoUsergrowthTaskConfigGetTopResult() *TaobaoUsergrowthTaskConfigGetTopResult {
	return poolTaobaoUsergrowthTaskConfigGetTopResult.Get().(*TaobaoUsergrowthTaskConfigGetTopResult)
}

// ReleaseTaobaoUsergrowthTaskConfigGetTopResult 释放TaobaoUsergrowthTaskConfigGetTopResult
func ReleaseTaobaoUsergrowthTaskConfigGetTopResult(v *TaobaoUsergrowthTaskConfigGetTopResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = nil
	v.Success = false
	poolTaobaoUsergrowthTaskConfigGetTopResult.Put(v)
}
