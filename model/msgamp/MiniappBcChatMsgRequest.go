package msgamp

import (
	"sync"
)

// MiniappBcChatMsgRequest 结构体
type MiniappBcChatMsgRequest struct {
	// 消息实例ID，在控制台申请到
	MsgInstanceId string `json:"msg_instance_id,omitempty" xml:"msg_instance_id,omitempty"`
	// 对应着商家C端APPID
	AppId string `json:"app_id,omitempty" xml:"app_id,omitempty"`
	// 消息自定义消息占位符
	MsgData string `json:"msg_data,omitempty" xml:"msg_data,omitempty"`
	// 自定义参数
	UrlParams string `json:"url_params,omitempty" xml:"url_params,omitempty"`
}

var poolMiniappBcChatMsgRequest = sync.Pool{
	New: func() any {
		return new(MiniappBcChatMsgRequest)
	},
}

// GetMiniappBcChatMsgRequest() 从对象池中获取MiniappBcChatMsgRequest
func GetMiniappBcChatMsgRequest() *MiniappBcChatMsgRequest {
	return poolMiniappBcChatMsgRequest.Get().(*MiniappBcChatMsgRequest)
}

// ReleaseMiniappBcChatMsgRequest 释放MiniappBcChatMsgRequest
func ReleaseMiniappBcChatMsgRequest(v *MiniappBcChatMsgRequest) {
	v.MsgInstanceId = ""
	v.AppId = ""
	v.MsgData = ""
	v.UrlParams = ""
	poolMiniappBcChatMsgRequest.Put(v)
}
