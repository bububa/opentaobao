package alimember

// SyncLevelSettingDto 结构体
type SyncLevelSettingDto struct {
	// main主账号 sub子账号
	UidType string `json:"uid_type,omitempty" xml:"uid_type,omitempty"`
	// level_type=2时，必须传递；
	LevelToGrowth string `json:"level_to_growth,omitempty" xml:"level_to_growth,omitempty"`
	// 1行为等级 2成长值等级
	LevelType int64 `json:"level_type,omitempty" xml:"level_type,omitempty"`
}
