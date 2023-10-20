package lsttrade

import (
	"sync"
)

// TopOrderChange2BrandownerDto 结构体
type TopOrderChange2BrandownerDto struct {
	// FORWARD_ORDER 正向订单表示正常购买流程,REVERSE_ORDER 逆向订单表示退款流程
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 退款单id
	RefundId int64 `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	// 主订单id
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
	// 子订单id
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 是否新建，包括订单新建和退款单新建
	NewCreated bool `json:"new_created,omitempty" xml:"new_created,omitempty"`
}

var poolTopOrderChange2BrandownerDto = sync.Pool{
	New: func() any {
		return new(TopOrderChange2BrandownerDto)
	},
}

// GetTopOrderChange2BrandownerDto() 从对象池中获取TopOrderChange2BrandownerDto
func GetTopOrderChange2BrandownerDto() *TopOrderChange2BrandownerDto {
	return poolTopOrderChange2BrandownerDto.Get().(*TopOrderChange2BrandownerDto)
}

// ReleaseTopOrderChange2BrandownerDto 释放TopOrderChange2BrandownerDto
func ReleaseTopOrderChange2BrandownerDto(v *TopOrderChange2BrandownerDto) {
	v.BizType = ""
	v.RefundId = 0
	v.MainOrderId = 0
	v.SubOrderId = 0
	v.NewCreated = false
	poolTopOrderChange2BrandownerDto.Put(v)
}
