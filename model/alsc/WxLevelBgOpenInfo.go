package alsc

import (
	"sync"
)

// WxLevelBgOpenInfo 结构体
type WxLevelBgOpenInfo struct {
	// 卡面URL
	BgUrl string `json:"bg_url,omitempty" xml:"bg_url,omitempty"`
	// 等级ID
	LevelId string `json:"level_id,omitempty" xml:"level_id,omitempty"`
	// 等级名称
	LevelName string `json:"level_name,omitempty" xml:"level_name,omitempty"`
}

var poolWxLevelBgOpenInfo = sync.Pool{
	New: func() any {
		return new(WxLevelBgOpenInfo)
	},
}

// GetWxLevelBgOpenInfo() 从对象池中获取WxLevelBgOpenInfo
func GetWxLevelBgOpenInfo() *WxLevelBgOpenInfo {
	return poolWxLevelBgOpenInfo.Get().(*WxLevelBgOpenInfo)
}

// ReleaseWxLevelBgOpenInfo 释放WxLevelBgOpenInfo
func ReleaseWxLevelBgOpenInfo(v *WxLevelBgOpenInfo) {
	v.BgUrl = ""
	v.LevelId = ""
	v.LevelName = ""
	poolWxLevelBgOpenInfo.Put(v)
}
