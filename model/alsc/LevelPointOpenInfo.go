package alsc

import (
	"sync"
)

// LevelPointOpenInfo 结构体
type LevelPointOpenInfo struct {
	// 等级ID
	LevelId string `json:"level_id,omitempty" xml:"level_id,omitempty"`
	// 等级名称
	LevelName string `json:"level_name,omitempty" xml:"level_name,omitempty"`
	// 积分奖励倍数
	Times string `json:"times,omitempty" xml:"times,omitempty"`
	// 是否已删除
	Deleted bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// 是否参与积分奖励
	UseSwitch bool `json:"use_switch,omitempty" xml:"use_switch,omitempty"`
}

var poolLevelPointOpenInfo = sync.Pool{
	New: func() any {
		return new(LevelPointOpenInfo)
	},
}

// GetLevelPointOpenInfo() 从对象池中获取LevelPointOpenInfo
func GetLevelPointOpenInfo() *LevelPointOpenInfo {
	return poolLevelPointOpenInfo.Get().(*LevelPointOpenInfo)
}

// ReleaseLevelPointOpenInfo 释放LevelPointOpenInfo
func ReleaseLevelPointOpenInfo(v *LevelPointOpenInfo) {
	v.LevelId = ""
	v.LevelName = ""
	v.Times = ""
	v.Deleted = false
	v.UseSwitch = false
	poolLevelPointOpenInfo.Put(v)
}
