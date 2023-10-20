package examination

import (
	"sync"
)

// Reportdiagnosemessagelist 结构体
type Reportdiagnosemessagelist struct {
	// 发送时间
	SendTime string `json:"send_time,omitempty" xml:"send_time,omitempty"`
	// 消息媒体链接
	MessageMediaUrl string `json:"message_media_url,omitempty" xml:"message_media_url,omitempty"`
	// 消息类型
	MessageType string `json:"message_type,omitempty" xml:"message_type,omitempty"`
	// 消息文字内容
	MessageContent string `json:"message_content,omitempty" xml:"message_content,omitempty"`
	// 接收者名称
	ReceiverName string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
	// 接收者ID
	ReceiverId string `json:"receiver_id,omitempty" xml:"receiver_id,omitempty"`
	// 发送者名称
	SenderName string `json:"sender_name,omitempty" xml:"sender_name,omitempty"`
	// 发送者ID
	SenderId string `json:"sender_id,omitempty" xml:"sender_id,omitempty"`
	// 序列号
	SeqNo string `json:"seq_no,omitempty" xml:"seq_no,omitempty"`
}

var poolReportdiagnosemessagelist = sync.Pool{
	New: func() any {
		return new(Reportdiagnosemessagelist)
	},
}

// GetReportdiagnosemessagelist() 从对象池中获取Reportdiagnosemessagelist
func GetReportdiagnosemessagelist() *Reportdiagnosemessagelist {
	return poolReportdiagnosemessagelist.Get().(*Reportdiagnosemessagelist)
}

// ReleaseReportdiagnosemessagelist 释放Reportdiagnosemessagelist
func ReleaseReportdiagnosemessagelist(v *Reportdiagnosemessagelist) {
	v.SendTime = ""
	v.MessageMediaUrl = ""
	v.MessageType = ""
	v.MessageContent = ""
	v.ReceiverName = ""
	v.ReceiverId = ""
	v.SenderName = ""
	v.SenderId = ""
	v.SeqNo = ""
	poolReportdiagnosemessagelist.Put(v)
}
