package icbudropshipping

// OrderPayRequest 结构体
type OrderPayRequest struct {
	// Order numbers to be paid，max size is  10
	OrderIdList []int64 `json:"order_id_list,omitempty" xml:"order_id_list>int64,omitempty"`
	// HTTP request Header "accept-language"
	AcceptLanguage string `json:"accept_language,omitempty" xml:"accept_language,omitempty"`
	// Screen resolution of the device used by the buyer
	ScreenResolution string `json:"screen_resolution,omitempty" xml:"screen_resolution,omitempty"`
	// HTTP request header "User-Agent"
	UserAgent string `json:"user_agent,omitempty" xml:"user_agent,omitempty"`
	// Buyer's IP
	UserIp string `json:"user_ip,omitempty" xml:"user_ip,omitempty"`
	// The time stamp of the buyer's first registration as dropper
	IsvDropShipperRegistrationTime int64 `json:"isv_drop_shipper_registration_time,omitempty" xml:"isv_drop_shipper_registration_time,omitempty"`
	// is PC ? true/false, current only support pc
	IsPc bool `json:"is_pc,omitempty" xml:"is_pc,omitempty"`
}
