package alsc

// CheckPayPasswdReq 结构体
type CheckPayPasswdReq struct {
	// 品牌ID 、 外部品牌id  2选1
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 密码
	PayPasswd string `json:"pay_passwd,omitempty" xml:"pay_passwd,omitempty"`
	// 顾客id
	CustomerId string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
	// 外部品牌id
	OutBrandId string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
	// 传入的密码是否已经加密
	Encrypted bool `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
}
