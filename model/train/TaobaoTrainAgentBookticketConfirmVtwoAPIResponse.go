package train

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentBookticketConfirmVtwoAPIResponse 火车票代理商接口——确认出票是否成功v2--增加鉴权校验 API返回值
// taobao.train.agent.bookticket.confirm.vtwo
//
// 火车票代理商接口——确认出票是否成功
type TaobaoTrainAgentBookticketConfirmVtwoAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentBookticketConfirmVtwoAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTrainAgentBookticketConfirmVtwoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTrainAgentBookticketConfirmVtwoAPIResponseModel).Reset()
}

// TaobaoTrainAgentBookticketConfirmVtwoAPIResponseModel is 火车票代理商接口——确认出票是否成功v2--增加鉴权校验 成功返回结果
type TaobaoTrainAgentBookticketConfirmVtwoAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_bookticket_confirm_vtwo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTrainAgentBookticketConfirmVtwoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoTrainAgentBookticketConfirmVtwoAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTrainAgentBookticketConfirmVtwoAPIResponse)
	},
}

// GetTaobaoTrainAgentBookticketConfirmVtwoAPIResponse 从 sync.Pool 获取 TaobaoTrainAgentBookticketConfirmVtwoAPIResponse
func GetTaobaoTrainAgentBookticketConfirmVtwoAPIResponse() *TaobaoTrainAgentBookticketConfirmVtwoAPIResponse {
	return poolTaobaoTrainAgentBookticketConfirmVtwoAPIResponse.Get().(*TaobaoTrainAgentBookticketConfirmVtwoAPIResponse)
}

// ReleaseTaobaoTrainAgentBookticketConfirmVtwoAPIResponse 将 TaobaoTrainAgentBookticketConfirmVtwoAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTrainAgentBookticketConfirmVtwoAPIResponse(v *TaobaoTrainAgentBookticketConfirmVtwoAPIResponse) {
	v.Reset()
	poolTaobaoTrainAgentBookticketConfirmVtwoAPIResponse.Put(v)
}
