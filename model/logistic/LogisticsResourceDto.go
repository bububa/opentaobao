package logistic

// LogisticsResourceDto 结构体
type LogisticsResourceDto struct {
	// 运单号校验正则表达式
	RegMailNo string `json:"reg_mail_no,omitempty" xml:"reg_mail_no,omitempty"`
	// 快递资源编码
	ResourceCode string `json:"resource_code,omitempty" xml:"resource_code,omitempty"`
	// 快递资源名称
	ResourceName string `json:"resource_name,omitempty" xml:"resource_name,omitempty"`
	// 快递公司id
	CompanyId int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
}
