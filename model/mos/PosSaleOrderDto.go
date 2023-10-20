package mos

import (
	"sync"
)

// PosSaleOrderDto 结构体
type PosSaleOrderDto struct {
	// 商品列表
	SaleItems []PosOrderSaleItemDto `json:"sale_items,omitempty" xml:"sale_items>pos_order_sale_item_dto,omitempty"`
	// 扩展信息
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
	// 会员卡号
	MemberCardNo string `json:"member_card_no,omitempty" xml:"member_card_no,omitempty"`
	// 会员电话
	MemberMobile string `json:"member_mobile,omitempty" xml:"member_mobile,omitempty"`
	// 操作员
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 开票单号
	SaleTicketNo string `json:"sale_ticket_no,omitempty" xml:"sale_ticket_no,omitempty"`
}

var poolPosSaleOrderDto = sync.Pool{
	New: func() any {
		return new(PosSaleOrderDto)
	},
}

// GetPosSaleOrderDto() 从对象池中获取PosSaleOrderDto
func GetPosSaleOrderDto() *PosSaleOrderDto {
	return poolPosSaleOrderDto.Get().(*PosSaleOrderDto)
}

// ReleasePosSaleOrderDto 释放PosSaleOrderDto
func ReleasePosSaleOrderDto(v *PosSaleOrderDto) {
	v.SaleItems = v.SaleItems[:0]
	v.ExtendParams = ""
	v.MemberCardNo = ""
	v.MemberMobile = ""
	v.Operator = ""
	v.SaleTicketNo = ""
	poolPosSaleOrderDto.Put(v)
}
