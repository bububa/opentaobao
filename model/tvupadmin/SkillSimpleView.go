package tvupadmin

import (
	"sync"
)

// SkillSimpleView 结构体
type SkillSimpleView struct {
	// 技能名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 图片地址
	Image string `json:"image,omitempty" xml:"image,omitempty"`
	// 是否在线
	IsOnline int64 `json:"is_online,omitempty" xml:"is_online,omitempty"`
	// bot技能关系表id
	BotSkillId int64 `json:"bot_skill_id,omitempty" xml:"bot_skill_id,omitempty"`
	// 技能id
	SkillId int64 `json:"skill_id,omitempty" xml:"skill_id,omitempty"`
	// 上下架
	DeleteToken int64 `json:"delete_token,omitempty" xml:"delete_token,omitempty"`
}

var poolSkillSimpleView = sync.Pool{
	New: func() any {
		return new(SkillSimpleView)
	},
}

// GetSkillSimpleView() 从对象池中获取SkillSimpleView
func GetSkillSimpleView() *SkillSimpleView {
	return poolSkillSimpleView.Get().(*SkillSimpleView)
}

// ReleaseSkillSimpleView 释放SkillSimpleView
func ReleaseSkillSimpleView(v *SkillSimpleView) {
	v.Name = ""
	v.Image = ""
	v.IsOnline = 0
	v.BotSkillId = 0
	v.SkillId = 0
	v.DeleteToken = 0
	poolSkillSimpleView.Put(v)
}
