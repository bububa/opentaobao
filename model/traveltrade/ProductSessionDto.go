package traveltrade

import (
	"sync"
)

// ProductSessionDto 结构体
type ProductSessionDto struct {
	// 开始时间。HH:mm
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 场次ID
	SessionId string `json:"session_id,omitempty" xml:"session_id,omitempty"`
	// 结束时间。HH:mm
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 场次库存
	Stock int64 `json:"stock,omitempty" xml:"stock,omitempty"`
	// 产品场次结算单价。单位：分
	WholesalePrice int64 `json:"wholesale_price,omitempty" xml:"wholesale_price,omitempty"`
	// 产品场次销售单价。单位：分
	RetailPrice int64 `json:"retail_price,omitempty" xml:"retail_price,omitempty"`
}

var poolProductSessionDto = sync.Pool{
	New: func() any {
		return new(ProductSessionDto)
	},
}

// GetProductSessionDto() 从对象池中获取ProductSessionDto
func GetProductSessionDto() *ProductSessionDto {
	return poolProductSessionDto.Get().(*ProductSessionDto)
}

// ReleaseProductSessionDto 释放ProductSessionDto
func ReleaseProductSessionDto(v *ProductSessionDto) {
	v.StartTime = ""
	v.SessionId = ""
	v.EndTime = ""
	v.Stock = 0
	v.WholesalePrice = 0
	v.RetailPrice = 0
	poolProductSessionDto.Put(v)
}
