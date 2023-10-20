package mtopopen

import (
	"sync"
)

// MsgSendResponse 结构体
type MsgSendResponse struct {
	// 短信发送失败结果码(成功情况无需关注)
	SmsErrorCode string `json:"sms_error_code,omitempty" xml:"sms_error_code,omitempty"`
	// 交易物流消息发送失败结果码(成功情况无需关注)
	MsgPushErrorCode string `json:"msg_push_error_code,omitempty" xml:"msg_push_error_code,omitempty"`
	// 发送操作执行时间
	OperateTime int64 `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// 短信发送是否成功
	SmsSuccess bool `json:"sms_success,omitempty" xml:"sms_success,omitempty"`
	// 交易物流消息是否发送成功
	MsgPushSuccess bool `json:"msg_push_success,omitempty" xml:"msg_push_success,omitempty"`
}

var poolMsgSendResponse = sync.Pool{
	New: func() any {
		return new(MsgSendResponse)
	},
}

// GetMsgSendResponse() 从对象池中获取MsgSendResponse
func GetMsgSendResponse() *MsgSendResponse {
	return poolMsgSendResponse.Get().(*MsgSendResponse)
}

// ReleaseMsgSendResponse 释放MsgSendResponse
func ReleaseMsgSendResponse(v *MsgSendResponse) {
	v.SmsErrorCode = ""
	v.MsgPushErrorCode = ""
	v.OperateTime = 0
	v.SmsSuccess = false
	v.MsgPushSuccess = false
	poolMsgSendResponse.Put(v)
}
