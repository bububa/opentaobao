package waybill

// WaybillOrderConfirmWaybillInfo 结构体
type WaybillOrderConfirmWaybillInfo struct {
	// 面单号
	WaybillCode string `json:"waybill_code,omitempty" xml:"waybill_code,omitempty"`
	// 包裹高，单位厘米
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
	// 包裹长，单位厘米
	Length int64 `json:"length,omitempty" xml:"length,omitempty"`
	// 体积, 单位 ml
	Volume int64 `json:"volume,omitempty" xml:"volume,omitempty"`
	// 重量,单位 g
	Weight int64 `json:"weight,omitempty" xml:"weight,omitempty"`
	// 包裹宽，单位厘米
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
}
