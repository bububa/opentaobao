package simba

// CrowdDto 结构体
type CrowdDto struct {
	// 用户所选择的人群名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 业务类型
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 用户在直通车的ID
	CustId int64 `json:"cust_id,omitempty" xml:"cust_id,omitempty"`
	// 人群包模版类型
	TemplateId int64 `json:"template_id,omitempty" xml:"template_id,omitempty"`
	// 人群包类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 人群id ,报表用
	CrowdId int64 `json:"crowd_id,omitempty" xml:"crowd_id,omitempty"`
	// 人群建议溢价,取值范围[5,300]
	FitDiscount int64 `json:"fit_discount,omitempty" xml:"fit_discount,omitempty"`
	// 所选择的人群id
	DmpcrowdId int64 `json:"dmpcrowd_id,omitempty" xml:"dmpcrowd_id,omitempty"`
}
