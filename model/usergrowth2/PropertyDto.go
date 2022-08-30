package usergrowth2

// PropertyDto 结构体
type PropertyDto struct {
	// 资产id
	PropertyId string `json:"property_id,omitempty" xml:"property_id,omitempty"`
	// 资产名称
	PropertyName string `json:"property_name,omitempty" xml:"property_name,omitempty"`
	// 资产业务类型
	PropertyBizType string `json:"property_biz_type,omitempty" xml:"property_biz_type,omitempty"`
	// 资产值
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}
