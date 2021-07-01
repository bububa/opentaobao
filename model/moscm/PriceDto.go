package moscm

// PriceDto 结构体
type PriceDto struct {
	// 专柜id(counter_id和out_counter_id必填一个，如果都填以counter_id为准)
	CounterId string `json:"counter_id,omitempty" xml:"counter_id,omitempty"`
	// 外部专柜id(经过配置后，可以传供应自己的专柜id)
	OutCounterId string `json:"out_counter_id,omitempty" xml:"out_counter_id,omitempty"`
	// 供应商商品Id
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 商品标准价
	StandardPrice string `json:"standard_price,omitempty" xml:"standard_price,omitempty"`
}
