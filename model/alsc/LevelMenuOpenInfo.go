package alsc

import (
	"sync"
)

// LevelMenuOpenInfo 结构体
type LevelMenuOpenInfo struct {
	// 等级ID
	LevelId string `json:"level_id,omitempty" xml:"level_id,omitempty"`
	// 等级名称
	LevelName string `json:"level_name,omitempty" xml:"level_name,omitempty"`
	// 菜品ID
	MenuId string `json:"menu_id,omitempty" xml:"menu_id,omitempty"`
	// 菜品名称
	MenuName string `json:"menu_name,omitempty" xml:"menu_name,omitempty"`
	// 是否已删除
	Deleted bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// 是否享受会员价
	UseSwitch bool `json:"use_switch,omitempty" xml:"use_switch,omitempty"`
}

var poolLevelMenuOpenInfo = sync.Pool{
	New: func() any {
		return new(LevelMenuOpenInfo)
	},
}

// GetLevelMenuOpenInfo() 从对象池中获取LevelMenuOpenInfo
func GetLevelMenuOpenInfo() *LevelMenuOpenInfo {
	return poolLevelMenuOpenInfo.Get().(*LevelMenuOpenInfo)
}

// ReleaseLevelMenuOpenInfo 释放LevelMenuOpenInfo
func ReleaseLevelMenuOpenInfo(v *LevelMenuOpenInfo) {
	v.LevelId = ""
	v.LevelName = ""
	v.MenuId = ""
	v.MenuName = ""
	v.Deleted = false
	v.UseSwitch = false
	poolLevelMenuOpenInfo.Put(v)
}
