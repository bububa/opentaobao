package scs

import (
	"sync"
)

// TaobaoOnebpDkxCrowdCrowdCoverageResultDto 结构体
type TaobaoOnebpDkxCrowdCrowdCoverageResultDto struct {
	// 返回消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
	// 结果总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 返回状态码
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoOnebpDkxCrowdCrowdCoverageResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoOnebpDkxCrowdCrowdCoverageResultDto)
	},
}

// GetTaobaoOnebpDkxCrowdCrowdCoverageResultDto() 从对象池中获取TaobaoOnebpDkxCrowdCrowdCoverageResultDto
func GetTaobaoOnebpDkxCrowdCrowdCoverageResultDto() *TaobaoOnebpDkxCrowdCrowdCoverageResultDto {
	return poolTaobaoOnebpDkxCrowdCrowdCoverageResultDto.Get().(*TaobaoOnebpDkxCrowdCrowdCoverageResultDto)
}

// ReleaseTaobaoOnebpDkxCrowdCrowdCoverageResultDto 释放TaobaoOnebpDkxCrowdCrowdCoverageResultDto
func ReleaseTaobaoOnebpDkxCrowdCrowdCoverageResultDto(v *TaobaoOnebpDkxCrowdCrowdCoverageResultDto) {
	v.Message = ""
	v.Result = 0
	v.TotalCount = 0
	v.ResultCode = nil
	v.Success = false
	poolTaobaoOnebpDkxCrowdCrowdCoverageResultDto.Put(v)
}
