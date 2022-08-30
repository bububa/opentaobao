package bus

// TvmContactInfo 结构体
type TvmContactInfo struct {
	// 下单人手机号，非特殊业务必填，影响商务结算
	MobileNumber string `json:"mobile_number,omitempty" xml:"mobile_number,omitempty"`
	// 下单人姓名
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
}
