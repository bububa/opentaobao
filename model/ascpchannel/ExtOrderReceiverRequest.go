package ascpchannel

// ExtOrderReceiverRequest 结构体
type ExtOrderReceiverRequest struct {
	// 收货人名称
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 收货人地址
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 地址编码
	DivisionId int64 `json:"division_id,omitempty" xml:"division_id,omitempty"`
}
