package wdk

import (
	"sync"
)

// QueryErpBillDto 结构体
type QueryErpBillDto struct {
	// 单据创建的开始时间点
	BeginDate string `json:"begin_date,omitempty" xml:"begin_date,omitempty"`
	// 单据创建的结束时间点
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 店仓code，指的是查询对象，对应一个物理店或仓编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 第几页
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 每页条数(最大1024)
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 单据类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolQueryErpBillDto = sync.Pool{
	New: func() any {
		return new(QueryErpBillDto)
	},
}

// GetQueryErpBillDto() 从对象池中获取QueryErpBillDto
func GetQueryErpBillDto() *QueryErpBillDto {
	return poolQueryErpBillDto.Get().(*QueryErpBillDto)
}

// ReleaseQueryErpBillDto 释放QueryErpBillDto
func ReleaseQueryErpBillDto(v *QueryErpBillDto) {
	v.BeginDate = ""
	v.EndDate = ""
	v.WarehouseCode = ""
	v.PageNo = 0
	v.PageSize = 0
	v.Type = 0
	poolQueryErpBillDto.Put(v)
}
