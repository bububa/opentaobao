package deliveryvoucher

// DeliveryVoucherDto 结构体
type DeliveryVoucherDto struct {
	// 券ID
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 券号
	No string `json:"no,omitempty" xml:"no,omitempty"`
}
