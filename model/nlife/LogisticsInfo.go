package nlife

// LogisticsInfo 结构体
type LogisticsInfo struct {
	// 货流详细信息
	LogisticsInfoDetails []LogisticsInfoDetail `json:"logistics_info_details,omitempty" xml:"logistics_info_details>logistics_info_detail,omitempty"`
	// 收货人
	Receiver string `json:"receiver,omitempty" xml:"receiver,omitempty"`
	// 收货地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 收货人联系电话
	PhoneNo string `json:"phone_no,omitempty" xml:"phone_no,omitempty"`
}
