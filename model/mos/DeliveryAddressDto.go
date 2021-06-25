package mos

// DeliveryAddressDto 
type DeliveryAddressDto struct {

    // 省
    Province   string `json:"province,omitempty"`

    // 市
    City   string `json:"city,omitempty"`

    // 区
    District   string `json:"district,omitempty"`

    // 镇
    Town   string `json:"town,omitempty"`

    // 编码
    DivisionId   int64 `json:"division_id,omitempty"`

    // 详细信息
    DetailAddress   string `json:"detail_address,omitempty"`

    // 邮编
    ZipCode   string `json:"zip_code,omitempty"`

}
