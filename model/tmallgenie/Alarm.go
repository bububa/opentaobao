package tmallgenie

// Alarm 结构体
type Alarm struct {
	// 闹钟音乐，第一次响起来的音乐
	AlarmMusic *AlarmMusic `json:"alarm_music,omitempty" xml:"alarm_music,omitempty"`
	// 闹钟音乐列表
	AlarmMusics []AlarmMusic `json:"alarm_musics,omitempty" xml:"alarm_musics>alarm_music,omitempty"`
	// 调度信息
	ScheduleInfo *ScheduleInfo `json:"schedule_info,omitempty" xml:"schedule_info,omitempty"`
	// 提醒方式，如APP，音箱等，如果有多个，用逗号分隔
	AlertWays []string `json:"alert_ways,omitempty" xml:"alert_ways>string,omitempty"`
	// musicPre
	MusicPre string `json:"music_pre,omitempty" xml:"music_pre,omitempty"`
	// 铃声Url
	MusicUrl string `json:"music_url,omitempty" xml:"music_url,omitempty"`
	// 闹铃主题
	Topic string `json:"topic,omitempty" xml:"topic,omitempty"`
	// 闹铃声音类型
	RingType string `json:"ring_type,omitempty" xml:"ring_type,omitempty"`
	// memo状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// memo_ID
	MemoId int64 `json:"memo_id,omitempty" xml:"memo_id,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
}
