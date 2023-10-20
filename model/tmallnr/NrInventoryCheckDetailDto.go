package tmallnr

import (
	"sync"
)

// NrInventoryCheckDetailDto 结构体
type NrInventoryCheckDetailDto struct {
	// 商家的商品编码
	ScItemCode string `json:"sc_item_code,omitempty" xml:"sc_item_code,omitempty"`
	// 幂等值
	SubOrderId string `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 淘宝后端商品的id号
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
}

var poolNrInventoryCheckDetailDto = sync.Pool{
	New: func() any {
		return new(NrInventoryCheckDetailDto)
	},
}

// GetNrInventoryCheckDetailDto() 从对象池中获取NrInventoryCheckDetailDto
func GetNrInventoryCheckDetailDto() *NrInventoryCheckDetailDto {
	return poolNrInventoryCheckDetailDto.Get().(*NrInventoryCheckDetailDto)
}

// ReleaseNrInventoryCheckDetailDto 释放NrInventoryCheckDetailDto
func ReleaseNrInventoryCheckDetailDto(v *NrInventoryCheckDetailDto) {
	v.ScItemCode = ""
	v.SubOrderId = ""
	v.Quantity = 0
	v.ScItemId = 0
	poolNrInventoryCheckDetailDto.Put(v)
}
