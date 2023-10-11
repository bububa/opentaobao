package ascp

// AddressNameResults 结构体
type AddressNameResults struct {
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 街道
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 错误原因
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
}
