package alsc

// LevelRuleOpenInfo 结构体
type LevelRuleOpenInfo struct {
	// 等级规则Id
	LevelId string `json:"level_id,omitempty" xml:"level_id,omitempty"`
	// 等级规则名称
	LevelName string `json:"level_name,omitempty" xml:"level_name,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 更新时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 创建者
	CreateBy string `json:"create_by,omitempty" xml:"create_by,omitempty"`
	// 更新者
	UpdateBy string `json:"update_by,omitempty" xml:"update_by,omitempty"`
	// 扩展字段
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 等级规则序号
	LevelNo int64 `json:"level_no,omitempty" xml:"level_no,omitempty"`
	// 等级对应成长值门槛
	Threshold int64 `json:"threshold,omitempty" xml:"threshold,omitempty"`
	// 升级赠送积分
	PresentPoint int64 `json:"present_point,omitempty" xml:"present_point,omitempty"`
	// 逻辑删除标志
	Deleted bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
}
