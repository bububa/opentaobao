package tmallgenie

import (
	"sync"
)

// CloudReportParam 结构体
type CloudReportParam struct {
	// 设备状态或者事件Map组成的Json字符串
	Payload string `json:"payload,omitempty" xml:"payload,omitempty"`
	// 设备id
	DeviceId string `json:"device_id,omitempty" xml:"device_id,omitempty"`
	// 用户accessToken
	UserAccessToken string `json:"user_access_token,omitempty" xml:"user_access_token,omitempty"`
	// 天猫精灵授权给厂商的userId
	OpenUserId string `json:"open_user_id,omitempty" xml:"open_user_id,omitempty"`
	// 消息id，问题定位用
	MessageId string `json:"message_id,omitempty" xml:"message_id,omitempty"`
	// 保留业务字段
	Extension string `json:"extension,omitempty" xml:"extension,omitempty"`
	// 账号类型（1：token 2：openUserId）
	AccountType int64 `json:"account_type,omitempty" xml:"account_type,omitempty"`
	// 技能id
	SkillId int64 `json:"skill_id,omitempty" xml:"skill_id,omitempty"`
	// 上报时间戳(毫秒)
	TimeStamp int64 `json:"time_stamp,omitempty" xml:"time_stamp,omitempty"`
	// 协议版本（原有协议版本为1，新版协议为2）
	PayloadVersion int64 `json:"payload_version,omitempty" xml:"payload_version,omitempty"`
	// 上报类型，1：属性上报 2：在离线上报 3：事件上报
	ReportType int64 `json:"report_type,omitempty" xml:"report_type,omitempty"`
}

var poolCloudReportParam = sync.Pool{
	New: func() any {
		return new(CloudReportParam)
	},
}

// GetCloudReportParam() 从对象池中获取CloudReportParam
func GetCloudReportParam() *CloudReportParam {
	return poolCloudReportParam.Get().(*CloudReportParam)
}

// ReleaseCloudReportParam 释放CloudReportParam
func ReleaseCloudReportParam(v *CloudReportParam) {
	v.Payload = ""
	v.DeviceId = ""
	v.UserAccessToken = ""
	v.OpenUserId = ""
	v.MessageId = ""
	v.Extension = ""
	v.AccountType = 0
	v.SkillId = 0
	v.TimeStamp = 0
	v.PayloadVersion = 0
	v.ReportType = 0
	poolCloudReportParam.Put(v)
}
