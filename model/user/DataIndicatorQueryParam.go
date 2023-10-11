package user

// DataIndicatorQueryParam 结构体
type DataIndicatorQueryParam struct {
	// 指标
	Measure []string `json:"measure,omitempty" xml:"measure>string,omitempty"`
	// 过滤条件
	Filter string `json:"filter,omitempty" xml:"filter,omitempty"`
	// 日期类型
	DateType string `json:"date_type,omitempty" xml:"date_type,omitempty"`
	// 固定
	NameSpace string `json:"name_space,omitempty" xml:"name_space,omitempty"`
	// CGP绑定的数据银行品牌商ID
	SmartId int64 `json:"smart_id,omitempty" xml:"smart_id,omitempty"`
}
