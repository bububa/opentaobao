package train

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentOrderFailAPIResponse 出票失败 API返回值
// taobao.train.agent.order.fail
//
// 出票失败
type TaobaoTrainAgentOrderFailAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentOrderFailAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTrainAgentOrderFailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTrainAgentOrderFailAPIResponseModel).Reset()
}

// TaobaoTrainAgentOrderFailAPIResponseModel is 出票失败 成功返回结果
type TaobaoTrainAgentOrderFailAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_order_fail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// rs
	Result *TapResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTrainAgentOrderFailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTrainAgentOrderFailAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTrainAgentOrderFailAPIResponse)
	},
}

// GetTaobaoTrainAgentOrderFailAPIResponse 从 sync.Pool 获取 TaobaoTrainAgentOrderFailAPIResponse
func GetTaobaoTrainAgentOrderFailAPIResponse() *TaobaoTrainAgentOrderFailAPIResponse {
	return poolTaobaoTrainAgentOrderFailAPIResponse.Get().(*TaobaoTrainAgentOrderFailAPIResponse)
}

// ReleaseTaobaoTrainAgentOrderFailAPIResponse 将 TaobaoTrainAgentOrderFailAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTrainAgentOrderFailAPIResponse(v *TaobaoTrainAgentOrderFailAPIResponse) {
	v.Reset()
	poolTaobaoTrainAgentOrderFailAPIResponse.Put(v)
}
