package traveltrade

// NormalVisaLogisticsInfo 结构体
type NormalVisaLogisticsInfo struct {
	// 必填，物流号
	PostNumber string `json:"post_number,omitempty" xml:"post_number,omitempty"`
	// 必填，物流公司编码
	PostCompanyCode string `json:"post_company_code,omitempty" xml:"post_company_code,omitempty"`
	// 必填，物流公司名称
	PostCompanyName string `json:"post_company_name,omitempty" xml:"post_company_name,omitempty"`
	// 选填，物流联系人手机号(顺丰物流需要)
	ConcatPhone string `json:"concat_phone,omitempty" xml:"concat_phone,omitempty"`
}
