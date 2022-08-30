package idleisv

// LogisticsRespResult 结构体
type LogisticsRespResult struct {
	// 全部公司列表
	CompanyList []CompanyList `json:"company_list,omitempty" xml:"company_list>company_list,omitempty"`
	// 热门公司列表
	HotCompanyList []HotCompanyList `json:"hot_company_list,omitempty" xml:"hot_company_list>hot_company_list,omitempty"`
	// 快递公司总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}
