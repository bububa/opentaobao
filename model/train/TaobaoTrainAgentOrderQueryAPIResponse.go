package train

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentOrderQueryAPIResponse 订单详情查询 API返回值
// taobao.train.agent.order.query
//
// 订单详情查询接口
type TaobaoTrainAgentOrderQueryAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentOrderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTrainAgentOrderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTrainAgentOrderQueryAPIResponseModel).Reset()
}

// TaobaoTrainAgentOrderQueryAPIResponseModel is 订单详情查询 成功返回结果
type TaobaoTrainAgentOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// rs
	Result *TapResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTrainAgentOrderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTrainAgentOrderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTrainAgentOrderQueryAPIResponse)
	},
}

// GetTaobaoTrainAgentOrderQueryAPIResponse 从 sync.Pool 获取 TaobaoTrainAgentOrderQueryAPIResponse
func GetTaobaoTrainAgentOrderQueryAPIResponse() *TaobaoTrainAgentOrderQueryAPIResponse {
	return poolTaobaoTrainAgentOrderQueryAPIResponse.Get().(*TaobaoTrainAgentOrderQueryAPIResponse)
}

// ReleaseTaobaoTrainAgentOrderQueryAPIResponse 将 TaobaoTrainAgentOrderQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTrainAgentOrderQueryAPIResponse(v *TaobaoTrainAgentOrderQueryAPIResponse) {
	v.Reset()
	poolTaobaoTrainAgentOrderQueryAPIResponse.Put(v)
}
