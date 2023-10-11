package ascpchannel

// WaybillGenSender 结构体
type WaybillGenSender struct {
	// 详细地址
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 省
	SenderProvince string `json:"sender_province,omitempty" xml:"sender_province,omitempty"`
	// 市
	SenderCity string `json:"sender_city,omitempty" xml:"sender_city,omitempty"`
	// 区
	SenderArea string `json:"sender_area,omitempty" xml:"sender_area,omitempty"`
}
