package lstmarketing

import (
	"sync"
)

// LstTopOrderDto 结构体
type LstTopOrderDto struct {
	// 营销活动列表
	PromotionDtoList []Promotiondtolist `json:"promotion_dto_list,omitempty" xml:"promotion_dto_list>promotiondtolist,omitempty"`
	// 子订单id
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 主订单id
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
}

var poolLstTopOrderDto = sync.Pool{
	New: func() any {
		return new(LstTopOrderDto)
	},
}

// GetLstTopOrderDto() 从对象池中获取LstTopOrderDto
func GetLstTopOrderDto() *LstTopOrderDto {
	return poolLstTopOrderDto.Get().(*LstTopOrderDto)
}

// ReleaseLstTopOrderDto 释放LstTopOrderDto
func ReleaseLstTopOrderDto(v *LstTopOrderDto) {
	v.PromotionDtoList = v.PromotionDtoList[:0]
	v.SubOrderId = 0
	v.MainOrderId = 0
	poolLstTopOrderDto.Put(v)
}
