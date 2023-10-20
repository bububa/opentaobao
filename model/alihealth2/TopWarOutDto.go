package alihealth2

import (
	"sync"
)

// TopWarOutDto 结构体
type TopWarOutDto struct {
	// 商品
	GoodsList []TopGoodsDto `json:"goods_list,omitempty" xml:"goods_list>top_goods_dto,omitempty"`
	// 出库单号
	BillNo string `json:"bill_no,omitempty" xml:"bill_no,omitempty"`
	// 单据日期
	BizDate string `json:"biz_date,omitempty" xml:"biz_date,omitempty"`
	// 孔雀翎O2O订单ID
	CepOrderId int64 `json:"cep_order_id,omitempty" xml:"cep_order_id,omitempty"`
	// 淘宝sellerId
	TbSellerId int64 `json:"tb_seller_id,omitempty" xml:"tb_seller_id,omitempty"`
	// 魔方ID
	CubeId int64 `json:"cube_id,omitempty" xml:"cube_id,omitempty"`
}

var poolTopWarOutDto = sync.Pool{
	New: func() any {
		return new(TopWarOutDto)
	},
}

// GetTopWarOutDto() 从对象池中获取TopWarOutDto
func GetTopWarOutDto() *TopWarOutDto {
	return poolTopWarOutDto.Get().(*TopWarOutDto)
}

// ReleaseTopWarOutDto 释放TopWarOutDto
func ReleaseTopWarOutDto(v *TopWarOutDto) {
	v.GoodsList = v.GoodsList[:0]
	v.BillNo = ""
	v.BizDate = ""
	v.CepOrderId = 0
	v.TbSellerId = 0
	v.CubeId = 0
	poolTopWarOutDto.Put(v)
}
