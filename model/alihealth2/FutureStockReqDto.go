package alihealth2

// FutureStockReqDto 结构体
type FutureStockReqDto struct {
	// 请求明细，数量不能超过10
	Items []FutureItem `json:"items,omitempty" xml:"items>future_item,omitempty"`
	// 请求流水号, 长度不能超过64, 相同的请求流水会被幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 供应商id
	SupplierId int64 `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
}
