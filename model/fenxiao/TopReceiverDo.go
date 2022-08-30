package fenxiao

// TopReceiverDo 结构体
type TopReceiverDo struct {
	// 收货人全名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 固定电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 移动电话
	MobilePhone string `json:"mobile_phone,omitempty" xml:"mobile_phone,omitempty"`
	// 收货人的详细地址信息
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 区/县
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 收货人的城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 邮政编码
	Zip string `json:"zip,omitempty" xml:"zip,omitempty"`
	// 收货人所在省份
	State string `json:"state,omitempty" xml:"state,omitempty"`
	// 收件人ID (Open Addressee ID)，长度在128个字符之内。
	Oaid string `json:"oaid,omitempty" xml:"oaid,omitempty"`
}
