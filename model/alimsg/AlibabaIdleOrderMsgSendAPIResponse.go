package alimsg

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleOrderMsgSendAPIResponse 虚拟发货消息发送接口 API返回值
// alibaba.idle.order.msg.send
//
// 用户下单后服务商期望自动发货，该接口用于给用户发送文本消息，主要用于卡券类等虚拟商品场景
type AlibabaIdleOrderMsgSendAPIResponse struct {
	model.CommonResponse
	AlibabaIdleOrderMsgSendAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleOrderMsgSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleOrderMsgSendAPIResponseModel).Reset()
}

// AlibabaIdleOrderMsgSendAPIResponseModel is 虚拟发货消息发送接口 成功返回结果
type AlibabaIdleOrderMsgSendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_order_msg_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否发送成功
	SendSuccess bool `json:"send_success,omitempty" xml:"send_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleOrderMsgSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SendSuccess = false
}

var poolAlibabaIdleOrderMsgSendAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleOrderMsgSendAPIResponse)
	},
}

// GetAlibabaIdleOrderMsgSendAPIResponse 从 sync.Pool 获取 AlibabaIdleOrderMsgSendAPIResponse
func GetAlibabaIdleOrderMsgSendAPIResponse() *AlibabaIdleOrderMsgSendAPIResponse {
	return poolAlibabaIdleOrderMsgSendAPIResponse.Get().(*AlibabaIdleOrderMsgSendAPIResponse)
}

// ReleaseAlibabaIdleOrderMsgSendAPIResponse 将 AlibabaIdleOrderMsgSendAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleOrderMsgSendAPIResponse(v *AlibabaIdleOrderMsgSendAPIResponse) {
	v.Reset()
	poolAlibabaIdleOrderMsgSendAPIResponse.Put(v)
}
