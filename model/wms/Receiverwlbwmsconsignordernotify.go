package wms

// Receiverwlbwmsconsignordernotify 结构体
type Receiverwlbwmsconsignordernotify struct {
	// 收件方省份
	ReceiverProvince string `json:"receiver_province,omitempty" xml:"receiver_province,omitempty"`
	// 收件方城市
	ReceiverCity string `json:"receiver_city,omitempty" xml:"receiver_city,omitempty"`
	// 收件方区县
	ReceiverArea string `json:"receiver_area,omitempty" xml:"receiver_area,omitempty"`
	// 收件方街道
	ReceiverTown string `json:"receiver_town,omitempty" xml:"receiver_town,omitempty"`
	// 收件方地址
	ReceiverAddress string `json:"receiver_address,omitempty" xml:"receiver_address,omitempty"`
	// 收件方邮编
	ReceiverZipCode string `json:"receiver_zip_code,omitempty" xml:"receiver_zip_code,omitempty"`
	// 收件人名称
	ReceiverName string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
	// 收件人Nick
	ReceiverNick string `json:"receiver_nick,omitempty" xml:"receiver_nick,omitempty"`
	// 收件人手机
	ReceiverMobile string `json:"receiver_mobile,omitempty" xml:"receiver_mobile,omitempty"`
	// 收件人电话
	ReceiverPhone string `json:"receiver_phone,omitempty" xml:"receiver_phone,omitempty"`
}
