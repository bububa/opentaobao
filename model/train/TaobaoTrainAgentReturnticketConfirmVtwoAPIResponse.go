package train

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentReturnticketConfirmVtwoAPIResponse 退票通知 API返回值
// taobao.train.agent.returnticket.confirm.vtwo
//
// 火车票代理商接口——退票通知回调
type TaobaoTrainAgentReturnticketConfirmVtwoAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentReturnticketConfirmVtwoAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTrainAgentReturnticketConfirmVtwoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTrainAgentReturnticketConfirmVtwoAPIResponseModel).Reset()
}

// TaobaoTrainAgentReturnticketConfirmVtwoAPIResponseModel is 退票通知 成功返回结果
type TaobaoTrainAgentReturnticketConfirmVtwoAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_returnticket_confirm_vtwo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// resultCode
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// resultMsg
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTrainAgentReturnticketConfirmVtwoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.ResultMsg = ""
	m.IsSuccess = false
}

var poolTaobaoTrainAgentReturnticketConfirmVtwoAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTrainAgentReturnticketConfirmVtwoAPIResponse)
	},
}

// GetTaobaoTrainAgentReturnticketConfirmVtwoAPIResponse 从 sync.Pool 获取 TaobaoTrainAgentReturnticketConfirmVtwoAPIResponse
func GetTaobaoTrainAgentReturnticketConfirmVtwoAPIResponse() *TaobaoTrainAgentReturnticketConfirmVtwoAPIResponse {
	return poolTaobaoTrainAgentReturnticketConfirmVtwoAPIResponse.Get().(*TaobaoTrainAgentReturnticketConfirmVtwoAPIResponse)
}

// ReleaseTaobaoTrainAgentReturnticketConfirmVtwoAPIResponse 将 TaobaoTrainAgentReturnticketConfirmVtwoAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTrainAgentReturnticketConfirmVtwoAPIResponse(v *TaobaoTrainAgentReturnticketConfirmVtwoAPIResponse) {
	v.Reset()
	poolTaobaoTrainAgentReturnticketConfirmVtwoAPIResponse.Put(v)
}
