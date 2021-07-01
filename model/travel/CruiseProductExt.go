package travel

// CruiseProductExt 结构体
type CruiseProductExt struct {
	// 选填，邮轮相关小费包含选项。境内邮轮： 1-"船票" 2-"岸上观光费" 3-"导游" 4-"其他费用" ...... 国际邮轮 1-"船票" 2-"港务费、邮轮税费" 3-"岸上观光费" 4-"签证费用" 5-"小费" 6-"领队费" 7-"其他费用"
	ShipFeeInclude []string `json:"ship_fee_include,omitempty" xml:"ship_fee_include>string,omitempty"`
	// 必填，邮轮船名
	ShipName string `json:"ship_name,omitempty" xml:"ship_name,omitempty"`
	// 必填，邮轮下船地点
	ShipDown string `json:"ship_down,omitempty" xml:"ship_down,omitempty"`
	// 必填，邮轮上船地点
	ShipUp string `json:"ship_up,omitempty" xml:"ship_up,omitempty"`
	// 必填，邮轮线路
	CruiseLine string `json:"cruise_line,omitempty" xml:"cruise_line,omitempty"`
	// 必填，邮轮公司
	CruiseCompany string `json:"cruise_company,omitempty" xml:"cruise_company,omitempty"`
}
