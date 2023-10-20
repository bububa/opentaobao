package train

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentTicketStatusCallbackAPIResponse 代理商票状态查询回调 API返回值
// taobao.train.agent.ticket.status.callback
//
// 代理商票状态查询结果回调
type TaobaoTrainAgentTicketStatusCallbackAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentTicketStatusCallbackAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTrainAgentTicketStatusCallbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTrainAgentTicketStatusCallbackAPIResponseModel).Reset()
}

// TaobaoTrainAgentTicketStatusCallbackAPIResponseModel is 代理商票状态查询回调 成功返回结果
type TaobaoTrainAgentTicketStatusCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_ticket_status_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功调用
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTrainAgentTicketStatusCallbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.IsSuccess = false
}

var poolTaobaoTrainAgentTicketStatusCallbackAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTrainAgentTicketStatusCallbackAPIResponse)
	},
}

// GetTaobaoTrainAgentTicketStatusCallbackAPIResponse 从 sync.Pool 获取 TaobaoTrainAgentTicketStatusCallbackAPIResponse
func GetTaobaoTrainAgentTicketStatusCallbackAPIResponse() *TaobaoTrainAgentTicketStatusCallbackAPIResponse {
	return poolTaobaoTrainAgentTicketStatusCallbackAPIResponse.Get().(*TaobaoTrainAgentTicketStatusCallbackAPIResponse)
}

// ReleaseTaobaoTrainAgentTicketStatusCallbackAPIResponse 将 TaobaoTrainAgentTicketStatusCallbackAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTrainAgentTicketStatusCallbackAPIResponse(v *TaobaoTrainAgentTicketStatusCallbackAPIResponse) {
	v.Reset()
	poolTaobaoTrainAgentTicketStatusCallbackAPIResponse.Put(v)
}
