package tvupadmin

// SkillSimpleView 结构体
type SkillSimpleView struct {
	// 图片地址
	Image string `json:"image,omitempty" xml:"image,omitempty"`
	// 技能名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 技能id
	SkillId int64 `json:"skill_id,omitempty" xml:"skill_id,omitempty"`
	// bot技能关系表id
	BotSkillId int64 `json:"bot_skill_id,omitempty" xml:"bot_skill_id,omitempty"`
	// 是否在线
	IsOnline int64 `json:"is_online,omitempty" xml:"is_online,omitempty"`
	// 上下架
	DeleteToken int64 `json:"delete_token,omitempty" xml:"delete_token,omitempty"`
}
