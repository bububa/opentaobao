package tblogistics

// LogisticsCompany 结构体
type LogisticsCompany struct {
	// 物流公司代码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 物流公司简称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 运单号验证正则表达式
	RegMailNo string `json:"reg_mail_no,omitempty" xml:"reg_mail_no,omitempty"`
	// 物流公司标识
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
