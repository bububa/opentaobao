package ascp

// AddressIdResults 结构体
type AddressIdResults struct {
	// 错误的地址ID
	AddressId string `json:"address_id,omitempty" xml:"address_id,omitempty"`
	// 错误原因
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
}
