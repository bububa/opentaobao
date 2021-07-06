package mos

// DeliveryAddressDto 结构体
type DeliveryAddressDto struct {
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 区
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 镇
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 详细信息
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 邮编
	ZipCode string `json:"zip_code,omitempty" xml:"zip_code,omitempty"`
	// 编码
	DivisionId int64 `json:"division_id,omitempty" xml:"division_id,omitempty"`
}
