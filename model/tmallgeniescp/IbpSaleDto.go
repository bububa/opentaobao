package tmallgeniescp

import (
	"sync"
)

// IbpSaleDto 结构体
type IbpSaleDto struct {
	// 关键日期值
	KeyFigureDate string `json:"key_figure_date,omitempty" xml:"key_figure_date,omitempty"`
	// 渠道号
	ChannelId string `json:"channel_id,omitempty" xml:"channel_id,omitempty"`
	// 物料号
	MaterielCode string `json:"materiel_code,omitempty" xml:"materiel_code,omitempty"`
	// 扩展参数
	ExtendJson string `json:"extend_json,omitempty" xml:"extend_json,omitempty"`
	// 租户
	Tenant string `json:"tenant,omitempty" xml:"tenant,omitempty"`
	// 渠道类型
	ChannelType string `json:"channel_type,omitempty" xml:"channel_type,omitempty"`
	// 激活数量
	ActivationQuantity int64 `json:"activation_quantity,omitempty" xml:"activation_quantity,omitempty"`
	// 实际进货数量
	SellInQuantity int64 `json:"sell_in_quantity,omitempty" xml:"sell_in_quantity,omitempty"`
	// 实际出货数量
	SellOutQuantity int64 `json:"sell_out_quantity,omitempty" xml:"sell_out_quantity,omitempty"`
	// 出货数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolIbpSaleDto = sync.Pool{
	New: func() any {
		return new(IbpSaleDto)
	},
}

// GetIbpSaleDto() 从对象池中获取IbpSaleDto
func GetIbpSaleDto() *IbpSaleDto {
	return poolIbpSaleDto.Get().(*IbpSaleDto)
}

// ReleaseIbpSaleDto 释放IbpSaleDto
func ReleaseIbpSaleDto(v *IbpSaleDto) {
	v.KeyFigureDate = ""
	v.ChannelId = ""
	v.MaterielCode = ""
	v.ExtendJson = ""
	v.Tenant = ""
	v.ChannelType = ""
	v.ActivationQuantity = 0
	v.SellInQuantity = 0
	v.SellOutQuantity = 0
	v.Quantity = 0
	poolIbpSaleDto.Put(v)
}
