package travel

import (
	"sync"
)

// TcwlItemExt 结构体
type TcwlItemExt struct {
	// 集合地信息
	GatherPlaces []GatherPlaceInfo `json:"gather_places,omitempty" xml:"gather_places>gather_place_info,omitempty"`
	// 活动时间
	ActivityTime string `json:"activity_time,omitempty" xml:"activity_time,omitempty"`
	// 活动地点
	ActivityPlace string `json:"activity_place,omitempty" xml:"activity_place,omitempty"`
	// 组织者旺旺
	OrgWangwang string `json:"org_wangwang,omitempty" xml:"org_wangwang,omitempty"`
	// 组织者电话
	OrgTel string `json:"org_tel,omitempty" xml:"org_tel,omitempty"`
	// 组织者介绍
	OrgIntroduce string `json:"org_introduce,omitempty" xml:"org_introduce,omitempty"`
	// 组织者名称
	Organization string `json:"organization,omitempty" xml:"organization,omitempty"`
	// 主题玩法
	TcwlThemePlay string `json:"tcwl_theme_play,omitempty" xml:"tcwl_theme_play,omitempty"`
	// 活动强度。1：底，2：中，3：高
	ActivityStrength int64 `json:"activity_strength,omitempty" xml:"activity_strength,omitempty"`
	// 玩乐主题
	TcwlTheme int64 `json:"tcwl_theme,omitempty" xml:"tcwl_theme,omitempty"`
}

var poolTcwlItemExt = sync.Pool{
	New: func() any {
		return new(TcwlItemExt)
	},
}

// GetTcwlItemExt() 从对象池中获取TcwlItemExt
func GetTcwlItemExt() *TcwlItemExt {
	return poolTcwlItemExt.Get().(*TcwlItemExt)
}

// ReleaseTcwlItemExt 释放TcwlItemExt
func ReleaseTcwlItemExt(v *TcwlItemExt) {
	v.GatherPlaces = v.GatherPlaces[:0]
	v.ActivityTime = ""
	v.ActivityPlace = ""
	v.OrgWangwang = ""
	v.OrgTel = ""
	v.OrgIntroduce = ""
	v.Organization = ""
	v.TcwlThemePlay = ""
	v.ActivityStrength = 0
	v.TcwlTheme = 0
	poolTcwlItemExt.Put(v)
}
