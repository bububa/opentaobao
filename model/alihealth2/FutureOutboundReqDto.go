package alihealth2

// FutureOutboundReqDto 结构体
type FutureOutboundReqDto struct {
	// 请求明细
	Items []FutureOutboundItem `json:"items,omitempty" xml:"items>future_outbound_item,omitempty"`
	// 请求流水号, 长度不能超过64, 相同的请求流水会被幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 期货退货ID
	FutureReturnId string `json:"future_return_id,omitempty" xml:"future_return_id,omitempty"`
	// 仓库编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 供应商ID
	SupplierId int64 `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 是否最后一次出库
	LastOutbound bool `json:"last_outbound,omitempty" xml:"last_outbound,omitempty"`
}
