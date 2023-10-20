package train

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentOrderIgnoreAPIResponse 忽略订单 API返回值
// taobao.train.agent.order.ignore
//
// 忽略订单
type TaobaoTrainAgentOrderIgnoreAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentOrderIgnoreAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTrainAgentOrderIgnoreAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTrainAgentOrderIgnoreAPIResponseModel).Reset()
}

// TaobaoTrainAgentOrderIgnoreAPIResponseModel is 忽略订单 成功返回结果
type TaobaoTrainAgentOrderIgnoreAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_order_ignore_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// rs
	Result *TapResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTrainAgentOrderIgnoreAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTrainAgentOrderIgnoreAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTrainAgentOrderIgnoreAPIResponse)
	},
}

// GetTaobaoTrainAgentOrderIgnoreAPIResponse 从 sync.Pool 获取 TaobaoTrainAgentOrderIgnoreAPIResponse
func GetTaobaoTrainAgentOrderIgnoreAPIResponse() *TaobaoTrainAgentOrderIgnoreAPIResponse {
	return poolTaobaoTrainAgentOrderIgnoreAPIResponse.Get().(*TaobaoTrainAgentOrderIgnoreAPIResponse)
}

// ReleaseTaobaoTrainAgentOrderIgnoreAPIResponse 将 TaobaoTrainAgentOrderIgnoreAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTrainAgentOrderIgnoreAPIResponse(v *TaobaoTrainAgentOrderIgnoreAPIResponse) {
	v.Reset()
	poolTaobaoTrainAgentOrderIgnoreAPIResponse.Put(v)
}
