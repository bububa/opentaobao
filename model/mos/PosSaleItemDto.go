package mos

import (
	"sync"
)

// PosSaleItemDto 结构体
type PosSaleItemDto struct {
	// 扩展参数
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
	// 商品ID
	GoodsId string `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
	// 商品名称
	GoodsName string `json:"goods_name,omitempty" xml:"goods_name,omitempty"`
	// 商品类型：  N：老单品；Y：喵货；C：云单品；M：零售+
	ItemType string `json:"item_type,omitempty" xml:"item_type,omitempty"`
	// 门店号
	MallNo string `json:"mall_no,omitempty" xml:"mall_no,omitempty"`
	// 数量，可小数点
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 开票单号（）
	SaleTicketNo string `json:"sale_ticket_no,omitempty" xml:"sale_ticket_no,omitempty"`
	// 结算码（退款的时候必须有）
	SettleCode string `json:"settle_code,omitempty" xml:"settle_code,omitempty"`
	// 专柜号
	ShopNo string `json:"shop_no,omitempty" xml:"shop_no,omitempty"`
	// 唯一码
	UniqueCode string `json:"unique_code,omitempty" xml:"unique_code,omitempty"`
	// 折扣金额，单位：分
	DiscountAmount int64 `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	// 商品行号
	GoodsLineNo int64 `json:"goods_line_no,omitempty" xml:"goods_line_no,omitempty"`
	// 商品原价，单位：分
	OriPrice int64 `json:"ori_price,omitempty" xml:"ori_price,omitempty"`
	// 应付金额，单位：分
	PayAmount int64 `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
}

var poolPosSaleItemDto = sync.Pool{
	New: func() any {
		return new(PosSaleItemDto)
	},
}

// GetPosSaleItemDto() 从对象池中获取PosSaleItemDto
func GetPosSaleItemDto() *PosSaleItemDto {
	return poolPosSaleItemDto.Get().(*PosSaleItemDto)
}

// ReleasePosSaleItemDto 释放PosSaleItemDto
func ReleasePosSaleItemDto(v *PosSaleItemDto) {
	v.ExtendParams = ""
	v.GoodsId = ""
	v.GoodsName = ""
	v.ItemType = ""
	v.MallNo = ""
	v.Quantity = ""
	v.SaleTicketNo = ""
	v.SettleCode = ""
	v.ShopNo = ""
	v.UniqueCode = ""
	v.DiscountAmount = 0
	v.GoodsLineNo = 0
	v.OriPrice = 0
	v.PayAmount = 0
	poolPosSaleItemDto.Put(v)
}
