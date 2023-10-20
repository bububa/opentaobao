package logistic

import (
	"sync"
)

// TmsPhoneCallInfoDto 结构体
type TmsPhoneCallInfoDto struct {
	// 电联人员名称（小件员）
	PhoneCallOperatorName string `json:"phone_call_operator_name,omitempty" xml:"phone_call_operator_name,omitempty"`
	// 电联人员联系方式（小件员/官方号码）
	PhoneCallOperatorPhone string `json:"phone_call_operator_phone,omitempty" xml:"phone_call_operator_phone,omitempty"`
	// 消费者联系方式（消费者）
	PhoneCallConsumerPhone string `json:"phone_call_consumer_phone,omitempty" xml:"phone_call_consumer_phone,omitempty"`
	// 消费者电联需求描述
	PhoneCallResultRemark string `json:"phone_call_result_remark,omitempty" xml:"phone_call_result_remark,omitempty"`
	// NOT_CONNECT, 未拨通；CONNECTED，已接通； NOBODY_ANSWER，无人接听
	ConnectionStatus string `json:"connection_status,omitempty" xml:"connection_status,omitempty"`
	// 拨打时间，操作时间（YYYY-MM-DD HH:MM:SS）
	CallTime string `json:"call_time,omitempty" xml:"call_time,omitempty"`
	// 接通时间，操作时间（YYYY-MM-DD HH:MM:SS）
	ConnectTime string `json:"connect_time,omitempty" xml:"connect_time,omitempty"`
	// 挂断时间，操作时间（YYYY-MM-DD HH:MM:SS）
	HangUpTime string `json:"hang_up_time,omitempty" xml:"hang_up_time,omitempty"`
	// 电联属性，SMART 智能云呼；NORMAL, 普通
	PhoneCallType string `json:"phone_call_type,omitempty" xml:"phone_call_type,omitempty"`
	// 挂断状态，CUSTOMER_HAND_UP,消费者挂断；SENDER_HANG_UP,小件员挂断 OTHER，其他（无法识别）
	HangUpType string `json:"hang_up_type,omitempty" xml:"hang_up_type,omitempty"`
	// 电联语音文件，用于客诉场景定责使用
	PhoneChatUrl string `json:"phone_chat_url,omitempty" xml:"phone_chat_url,omitempty"`
	// 接通时长,  单位s
	ConnectTimeLength int64 `json:"connect_time_length,omitempty" xml:"connect_time_length,omitempty"`
}

var poolTmsPhoneCallInfoDto = sync.Pool{
	New: func() any {
		return new(TmsPhoneCallInfoDto)
	},
}

// GetTmsPhoneCallInfoDto() 从对象池中获取TmsPhoneCallInfoDto
func GetTmsPhoneCallInfoDto() *TmsPhoneCallInfoDto {
	return poolTmsPhoneCallInfoDto.Get().(*TmsPhoneCallInfoDto)
}

// ReleaseTmsPhoneCallInfoDto 释放TmsPhoneCallInfoDto
func ReleaseTmsPhoneCallInfoDto(v *TmsPhoneCallInfoDto) {
	v.PhoneCallOperatorName = ""
	v.PhoneCallOperatorPhone = ""
	v.PhoneCallConsumerPhone = ""
	v.PhoneCallResultRemark = ""
	v.ConnectionStatus = ""
	v.CallTime = ""
	v.ConnectTime = ""
	v.HangUpTime = ""
	v.PhoneCallType = ""
	v.HangUpType = ""
	v.PhoneChatUrl = ""
	v.ConnectTimeLength = 0
	poolTmsPhoneCallInfoDto.Put(v)
}
