package scs

import (
	"sync"
)

// TaobaoOnebpDkxAccountAccountFreezeResultDto 结构体
type TaobaoOnebpDkxAccountAccountFreezeResultDto struct {
	// 返回消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 结果总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 返回状态码
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

var poolTaobaoOnebpDkxAccountAccountFreezeResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoOnebpDkxAccountAccountFreezeResultDto)
	},
}

// GetTaobaoOnebpDkxAccountAccountFreezeResultDto() 从对象池中获取TaobaoOnebpDkxAccountAccountFreezeResultDto
func GetTaobaoOnebpDkxAccountAccountFreezeResultDto() *TaobaoOnebpDkxAccountAccountFreezeResultDto {
	return poolTaobaoOnebpDkxAccountAccountFreezeResultDto.Get().(*TaobaoOnebpDkxAccountAccountFreezeResultDto)
}

// ReleaseTaobaoOnebpDkxAccountAccountFreezeResultDto 释放TaobaoOnebpDkxAccountAccountFreezeResultDto
func ReleaseTaobaoOnebpDkxAccountAccountFreezeResultDto(v *TaobaoOnebpDkxAccountAccountFreezeResultDto) {
	v.Message = ""
	v.TotalCount = 0
	v.ResultCode = nil
	v.Success = false
	v.Result = false
	poolTaobaoOnebpDkxAccountAccountFreezeResultDto.Put(v)
}
