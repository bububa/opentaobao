package waybill

// AddressArea 结构体
type AddressArea struct {
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
}
