package fenxiao

// Receiver 结构体
type Receiver struct {
	// 邮政编码
	Zip string `json:"zip,omitempty" xml:"zip,omitempty"`
	// 移动电话
	MobilePhone string `json:"mobile_phone,omitempty" xml:"mobile_phone,omitempty"`
	// 固定电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 收货人的详细地址信息
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 收货人全名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 收货人所在省份
	State string `json:"state,omitempty" xml:"state,omitempty"`
	// 区/县
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 证件号
	CardId string `json:"card_id,omitempty" xml:"card_id,omitempty"`
	// 收货人的城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
}
