package waybill

// WaybillAddress 结构体
type WaybillAddress struct {
	// 区名称（三级地址）
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 一级地址（省、直辖市
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 街道\镇名称（四级地址）
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 详细地址
	AddressDetail string `json:"address_detail,omitempty" xml:"address_detail,omitempty"`
	// 市名称（二级地址）
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 末级地址
	DivisionId int64 `json:"division_id,omitempty" xml:"division_id,omitempty"`
	// waybill 地址记录ID(非地址库ID)
	WaybillAddressId int64 `json:"waybill_address_id,omitempty" xml:"waybill_address_id,omitempty"`
}
