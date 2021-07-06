package travel

// TcwlItemExt 结构体
type TcwlItemExt struct {
	// 选填，集合地信息。当玩乐主题为 户外旅行 时，集合地信息才选填
	GatherPlaces []GatherPlaceInfo `json:"gather_places,omitempty" xml:"gather_places>gather_place_info,omitempty"`
	// 选填，主题玩法。根据 玩乐主题 选择的不同，主题玩法只有特定的几种枚举值，具体可用的枚举值详见网页端发布商品时该字段所在下拉框的可选值
	TcwlThemePlay string `json:"tcwl_theme_play,omitempty" xml:"tcwl_theme_play,omitempty"`
	// 选填，组织者名称
	Organization string `json:"organization,omitempty" xml:"organization,omitempty"`
	// 选填，组织者介绍
	OrgIntroduce string `json:"org_introduce,omitempty" xml:"org_introduce,omitempty"`
	// 选填，组织者电话
	OrgTel string `json:"org_tel,omitempty" xml:"org_tel,omitempty"`
	// 选填，组织者旺旺
	OrgWangwang string `json:"org_wangwang,omitempty" xml:"org_wangwang,omitempty"`
	// 特殊必填，活动地点。当“玩乐主题”不是 户外旅行 活动时，活动地点，活动时间必填
	ActivityPlace string `json:"activity_place,omitempty" xml:"activity_place,omitempty"`
	// 特殊必填，活动时间。当“玩乐主题”不是 户外旅行 活动时，活动地点，活动时间必填
	ActivityTime string `json:"activity_time,omitempty" xml:"activity_time,omitempty"`
	// 必填，玩乐主题。 1:会展/市集展览旅行, 2:少儿兴趣/DIY/体验, 3:户外旅行, 4:约会/派对, 5:讲座沙龙, 6:运动赛事旅行, 7:音乐演出
	TcwlTheme int64 `json:"tcwl_theme,omitempty" xml:"tcwl_theme,omitempty"`
	// 选填，活动强度。1：底，2：中，3：高
	ActivityStrength int64 `json:"activity_strength,omitempty" xml:"activity_strength,omitempty"`
}
