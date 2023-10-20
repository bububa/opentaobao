package qianniu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuTaskMessageSendAPIResponse 发送任务提醒消息 API返回值
// taobao.qianniu.task.message.send
//
// 如果taskid不为空，则只发给task对应的单个接收人。如果taskid为空，则发给metadata_id对应的所有接收人。消息会以任务消息的形式发给客户端。
type TaobaoQianniuTaskMessageSendAPIResponse struct {
	model.CommonResponse
	TaobaoQianniuTaskMessageSendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQianniuTaskMessageSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQianniuTaskMessageSendAPIResponseModel).Reset()
}

// TaobaoQianniuTaskMessageSendAPIResponseModel is 发送任务提醒消息 成功返回结果
type TaobaoQianniuTaskMessageSendAPIResponseModel struct {
	XMLName xml.Name `xml:"qianniu_task_message_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQianniuTaskMessageSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolTaobaoQianniuTaskMessageSendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQianniuTaskMessageSendAPIResponse)
	},
}

// GetTaobaoQianniuTaskMessageSendAPIResponse 从 sync.Pool 获取 TaobaoQianniuTaskMessageSendAPIResponse
func GetTaobaoQianniuTaskMessageSendAPIResponse() *TaobaoQianniuTaskMessageSendAPIResponse {
	return poolTaobaoQianniuTaskMessageSendAPIResponse.Get().(*TaobaoQianniuTaskMessageSendAPIResponse)
}

// ReleaseTaobaoQianniuTaskMessageSendAPIResponse 将 TaobaoQianniuTaskMessageSendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQianniuTaskMessageSendAPIResponse(v *TaobaoQianniuTaskMessageSendAPIResponse) {
	v.Reset()
	poolTaobaoQianniuTaskMessageSendAPIResponse.Put(v)
}
