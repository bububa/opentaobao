package alitripmerchant

import (
	"sync"
)

// QueryLotteryDataDto 结构体
type QueryLotteryDataDto struct {
	// 查询开始时间
	OperationStartTime string `json:"operation_start_time,omitempty" xml:"operation_start_time,omitempty"`
	// 查询结束时间
	OperationEndTime string `json:"operation_end_time,omitempty" xml:"operation_end_time,omitempty"`
	// 操作类型
	OperationType string `json:"operation_type,omitempty" xml:"operation_type,omitempty"`
	// 页数
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 单页数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolQueryLotteryDataDto = sync.Pool{
	New: func() any {
		return new(QueryLotteryDataDto)
	},
}

// GetQueryLotteryDataDto() 从对象池中获取QueryLotteryDataDto
func GetQueryLotteryDataDto() *QueryLotteryDataDto {
	return poolQueryLotteryDataDto.Get().(*QueryLotteryDataDto)
}

// ReleaseQueryLotteryDataDto 释放QueryLotteryDataDto
func ReleaseQueryLotteryDataDto(v *QueryLotteryDataDto) {
	v.OperationStartTime = ""
	v.OperationEndTime = ""
	v.OperationType = ""
	v.PageNo = 0
	v.PageSize = 0
	poolQueryLotteryDataDto.Put(v)
}
