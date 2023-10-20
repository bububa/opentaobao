package tmc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcMessagesConsumeAPIResponse 消费多条消息 API返回值
// taobao.tmc.messages.consume
//
// 消费多条消息。消费时如果没有返回消息，建议做控制，不要一直调api，浪费应用的流量。如对程序做好优化，若没有消息则，sleep 100ms 等。
type TaobaoTmcMessagesConsumeAPIResponse struct {
	model.CommonResponse
	TaobaoTmcMessagesConsumeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTmcMessagesConsumeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTmcMessagesConsumeAPIResponseModel).Reset()
}

// TaobaoTmcMessagesConsumeAPIResponseModel is 消费多条消息 成功返回结果
type TaobaoTmcMessagesConsumeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmc_messages_consume_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 消息列表
	Messages []TmcMessage `json:"messages,omitempty" xml:"messages>tmc_message,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTmcMessagesConsumeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Messages = m.Messages[:0]
}

var poolTaobaoTmcMessagesConsumeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTmcMessagesConsumeAPIResponse)
	},
}

// GetTaobaoTmcMessagesConsumeAPIResponse 从 sync.Pool 获取 TaobaoTmcMessagesConsumeAPIResponse
func GetTaobaoTmcMessagesConsumeAPIResponse() *TaobaoTmcMessagesConsumeAPIResponse {
	return poolTaobaoTmcMessagesConsumeAPIResponse.Get().(*TaobaoTmcMessagesConsumeAPIResponse)
}

// ReleaseTaobaoTmcMessagesConsumeAPIResponse 将 TaobaoTmcMessagesConsumeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTmcMessagesConsumeAPIResponse(v *TaobaoTmcMessagesConsumeAPIResponse) {
	v.Reset()
	poolTaobaoTmcMessagesConsumeAPIResponse.Put(v)
}
