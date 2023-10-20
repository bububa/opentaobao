package train

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentOrderConfirmAPIResponse 确认出票 API返回值
// taobao.train.agent.order.confirm
//
// 确认出票
type TaobaoTrainAgentOrderConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentOrderConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTrainAgentOrderConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTrainAgentOrderConfirmAPIResponseModel).Reset()
}

// TaobaoTrainAgentOrderConfirmAPIResponseModel is 确认出票 成功返回结果
type TaobaoTrainAgentOrderConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_order_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// rs
	Result *TapResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTrainAgentOrderConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTrainAgentOrderConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTrainAgentOrderConfirmAPIResponse)
	},
}

// GetTaobaoTrainAgentOrderConfirmAPIResponse 从 sync.Pool 获取 TaobaoTrainAgentOrderConfirmAPIResponse
func GetTaobaoTrainAgentOrderConfirmAPIResponse() *TaobaoTrainAgentOrderConfirmAPIResponse {
	return poolTaobaoTrainAgentOrderConfirmAPIResponse.Get().(*TaobaoTrainAgentOrderConfirmAPIResponse)
}

// ReleaseTaobaoTrainAgentOrderConfirmAPIResponse 将 TaobaoTrainAgentOrderConfirmAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTrainAgentOrderConfirmAPIResponse(v *TaobaoTrainAgentOrderConfirmAPIResponse) {
	v.Reset()
	poolTaobaoTrainAgentOrderConfirmAPIResponse.Put(v)
}
