package tmallgenie

import (
	"sync"
)

// Meeting 结构体
type Meeting struct {
	// 会议必选参与人员
	RequiredParticipants []string `json:"required_participants,omitempty" xml:"required_participants>string,omitempty"`
	// 会议可选参与人员
	OptionalParticipants []string `json:"optional_participants,omitempty" xml:"optional_participants>string,omitempty"`
	// 提醒方式，如APP，音箱等，如果有多个，用逗号分隔
	AlertWays []string `json:"alert_ways,omitempty" xml:"alert_ways>string,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// memo状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 会议预计开始时间
	ExpectedStartTime string `json:"expected_start_time,omitempty" xml:"expected_start_time,omitempty"`
	// 会议预计结束时间
	ExpectedEndTime string `json:"expected_end_time,omitempty" xml:"expected_end_time,omitempty"`
	// 期望提醒时间
	ExpectedRemindTime string `json:"expected_remind_time,omitempty" xml:"expected_remind_time,omitempty"`
	// 会议地点
	Location string `json:"location,omitempty" xml:"location,omitempty"`
	// 会议组织者
	Organizer string `json:"organizer,omitempty" xml:"organizer,omitempty"`
	// 会议内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 会议主题
	Topic string `json:"topic,omitempty" xml:"topic,omitempty"`
	// 铃声Url
	MusicUrl string `json:"music_url,omitempty" xml:"music_url,omitempty"`
	// memo_ID
	MemoId int64 `json:"memo_id,omitempty" xml:"memo_id,omitempty"`
	// 调度信息
	ScheduleInfo *ScheduleInfo `json:"schedule_info,omitempty" xml:"schedule_info,omitempty"`
}

var poolMeeting = sync.Pool{
	New: func() any {
		return new(Meeting)
	},
}

// GetMeeting() 从对象池中获取Meeting
func GetMeeting() *Meeting {
	return poolMeeting.Get().(*Meeting)
}

// ReleaseMeeting 释放Meeting
func ReleaseMeeting(v *Meeting) {
	v.RequiredParticipants = v.RequiredParticipants[:0]
	v.OptionalParticipants = v.OptionalParticipants[:0]
	v.AlertWays = v.AlertWays[:0]
	v.GmtCreate = ""
	v.GmtModified = ""
	v.Uuid = ""
	v.Status = ""
	v.ExpectedStartTime = ""
	v.ExpectedEndTime = ""
	v.ExpectedRemindTime = ""
	v.Location = ""
	v.Organizer = ""
	v.Content = ""
	v.Topic = ""
	v.MusicUrl = ""
	v.MemoId = 0
	v.ScheduleInfo = nil
	poolMeeting.Put(v)
}
