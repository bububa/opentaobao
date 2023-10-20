package iot

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopmessagesendAPIResponse 发送留言 API返回值
// taobao.ailab.aicloud.top.message.send
//
// 供准入的外部用户实现发送留言功能，APP端发送，设备端读取
type TaobaoailabaicloudtopmessagesendAPIResponse struct {
	model.CommonResponse
	TaobaoailabaicloudtopmessagesendAPIResponseModel
}

// TaobaoailabaicloudtopmessagesendAPIResponseModel is 发送留言 成功返回结果
type TaobaoailabaicloudtopmessagesendAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_message_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// model
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}
