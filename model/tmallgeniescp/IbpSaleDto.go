package tmallgeniescp

// IbpSaleDto 结构体
type IbpSaleDto struct {
	// 渠道号
	ChannelId string `json:"channel_id,omitempty" xml:"channel_id,omitempty"`
	// 物料号
	MaterielCode string `json:"materiel_code,omitempty" xml:"materiel_code,omitempty"`
	// 扩展参数
	ExtendJson string `json:"extend_json,omitempty" xml:"extend_json,omitempty"`
	// 租户
	Tenant string `json:"tenant,omitempty" xml:"tenant,omitempty"`
	// 关键日期值
	KeyFigureDate string `json:"key_figure_date,omitempty" xml:"key_figure_date,omitempty"`
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
