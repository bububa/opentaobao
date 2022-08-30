package ascp

// HiErpReceiverDto 结构体
type HiErpReceiverDto struct {
	// 收件人姓名
	ReceiverName string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
	// 手机号码
	CellPhone string `json:"cell_phone,omitempty" xml:"cell_phone,omitempty"`
	// 电话号码
	TelePhone string `json:"tele_phone,omitempty" xml:"tele_phone,omitempty"`
	// 国家
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 街道
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 详细地址
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 扩展字段，json格式
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
}
