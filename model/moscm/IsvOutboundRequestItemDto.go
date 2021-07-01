package moscm

// IsvOutboundRequestItemDto 结构体
type IsvOutboundRequestItemDto struct {
	// 外部id
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
}
