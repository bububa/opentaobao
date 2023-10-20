package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMessageSendAPIResponse 发送留言 API返回值
// taobao.ailab.aicloud.top.message.send
//
// 供准入的外部用户实现发送留言功能，APP端发送，设备端读取
type TaobaoAilabAicloudTopMessageSendAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopMessageSendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopMessageSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopMessageSendAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopMessageSendAPIResponseModel is 发送留言 成功返回结果
type TaobaoAilabAicloudTopMessageSendAPIResponseModel struct {
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

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopMessageSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.IsSuccess = false
	m.Model = false
}

var poolTaobaoAilabAicloudTopMessageSendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopMessageSendAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopMessageSendAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopMessageSendAPIResponse
func GetTaobaoAilabAicloudTopMessageSendAPIResponse() *TaobaoAilabAicloudTopMessageSendAPIResponse {
	return poolTaobaoAilabAicloudTopMessageSendAPIResponse.Get().(*TaobaoAilabAicloudTopMessageSendAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopMessageSendAPIResponse 将 TaobaoAilabAicloudTopMessageSendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopMessageSendAPIResponse(v *TaobaoAilabAicloudTopMessageSendAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopMessageSendAPIResponse.Put(v)
}
