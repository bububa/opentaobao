package travel

// CruiseItemExt 结构体
type CruiseItemExt struct {
	// 邮轮费用包含
	ShipFeeInclude string `json:"ship_fee_include,omitempty" xml:"ship_fee_include,omitempty"`
	// 邮轮船名
	ShipName string `json:"ship_name,omitempty" xml:"ship_name,omitempty"`
	// 邮轮公司
	CruiseCompany string `json:"cruise_company,omitempty" xml:"cruise_company,omitempty"`
	// 邮轮线路
	CruiseLine string `json:"cruise_line,omitempty" xml:"cruise_line,omitempty"`
	// 下船地点
	ShipDown string `json:"ship_down,omitempty" xml:"ship_down,omitempty"`
	// 上船地点
	ShipUp string `json:"ship_up,omitempty" xml:"ship_up,omitempty"`
	// 邮轮数据版本
	CruiseItemVersion string `json:"cruise_item_version,omitempty" xml:"cruise_item_version,omitempty"`
	// 邮轮具体航线
	CruiseSubLine string `json:"cruise_sub_line,omitempty" xml:"cruise_sub_line,omitempty"`
}
