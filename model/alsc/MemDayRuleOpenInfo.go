package alsc

// MemDayRuleOpenInfo 结构体
type MemDayRuleOpenInfo struct {
	// 是否已删除
	Deleted bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 会员等级和特价菜单的关系
	LevelMenuList []LevelMenuOpenInfo `json:"level_menu_list,omitempty" xml:"level_menu_list>level_menu_open_info,omitempty"`
	// 会员日积分奖励
	LevelPointList []LevelPointOpenInfo `json:"level_point_list,omitempty" xml:"level_point_list>level_point_open_info,omitempty"`
	// 会员日是哪天：每周几，每月几天 逗号分隔的字符串
	MemDayCircle string `json:"mem_day_circle,omitempty" xml:"mem_day_circle,omitempty"`
	// 会员周期，WEEK:每周，MONTH:每月
	MemDayCircleType string `json:"mem_day_circle_type,omitempty" xml:"mem_day_circle_type,omitempty"`
	// memDayRuleId
	MemDayRuleId string `json:"mem_day_rule_id,omitempty" xml:"mem_day_rule_id,omitempty"`
	// 会员日期开关，总开关
	MemDaySwitch bool `json:"mem_day_switch,omitempty" xml:"mem_day_switch,omitempty"`
	// 会员日特价菜单开关
	MenuSwitch bool `json:"menu_switch,omitempty" xml:"menu_switch,omitempty"`
	// 积分奖励开关
	PointSwitch bool `json:"point_switch,omitempty" xml:"point_switch,omitempty"`
	// 扩展信息
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
}
