package trade

import (
	"sync"
)

// OrderInfoResultDto 结构体
type OrderInfoResultDto struct {
	// 订单信息list
	OrderInfoList []OrderInfoDto `json:"order_info_list,omitempty" xml:"order_info_list>order_info_dto,omitempty"`
	// 当前页
	CurPageNo int64 `json:"cur_page_no,omitempty" xml:"cur_page_no,omitempty"`
	// 每页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 是否下一页
	HasNextPage bool `json:"has_next_page,omitempty" xml:"has_next_page,omitempty"`
}

var poolOrderInfoResultDto = sync.Pool{
	New: func() any {
		return new(OrderInfoResultDto)
	},
}

// GetOrderInfoResultDto() 从对象池中获取OrderInfoResultDto
func GetOrderInfoResultDto() *OrderInfoResultDto {
	return poolOrderInfoResultDto.Get().(*OrderInfoResultDto)
}

// ReleaseOrderInfoResultDto 释放OrderInfoResultDto
func ReleaseOrderInfoResultDto(v *OrderInfoResultDto) {
	v.OrderInfoList = v.OrderInfoList[:0]
	v.CurPageNo = 0
	v.PageSize = 0
	v.HasNextPage = false
	poolOrderInfoResultDto.Put(v)
}
