package alsc

// WxLevelBgOpenInfo 结构体
type WxLevelBgOpenInfo struct {
	// 卡面URL
	BgUrl string `json:"bg_url,omitempty" xml:"bg_url,omitempty"`
	// 等级ID
	LevelId string `json:"level_id,omitempty" xml:"level_id,omitempty"`
	// 等级名称
	LevelName string `json:"level_name,omitempty" xml:"level_name,omitempty"`
}
