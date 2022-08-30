package util

// PriceTopDto 结构体
type PriceTopDto struct {
	// 溢价信息，范围5-300，不在范围内会默认设置为5
	Discount int64 `json:"discount,omitempty" xml:"discount,omitempty"`
}
