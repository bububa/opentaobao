package train

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentOrderLockAPIResponse 锁单 API返回值
// taobao.train.agent.order.lock
//
// 锁单
type TaobaoTrainAgentOrderLockAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentOrderLockAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTrainAgentOrderLockAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTrainAgentOrderLockAPIResponseModel).Reset()
}

// TaobaoTrainAgentOrderLockAPIResponseModel is 锁单 成功返回结果
type TaobaoTrainAgentOrderLockAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_order_lock_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// rs
	Result *TapResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTrainAgentOrderLockAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTrainAgentOrderLockAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTrainAgentOrderLockAPIResponse)
	},
}

// GetTaobaoTrainAgentOrderLockAPIResponse 从 sync.Pool 获取 TaobaoTrainAgentOrderLockAPIResponse
func GetTaobaoTrainAgentOrderLockAPIResponse() *TaobaoTrainAgentOrderLockAPIResponse {
	return poolTaobaoTrainAgentOrderLockAPIResponse.Get().(*TaobaoTrainAgentOrderLockAPIResponse)
}

// ReleaseTaobaoTrainAgentOrderLockAPIResponse 将 TaobaoTrainAgentOrderLockAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTrainAgentOrderLockAPIResponse(v *TaobaoTrainAgentOrderLockAPIResponse) {
	v.Reset()
	poolTaobaoTrainAgentOrderLockAPIResponse.Put(v)
}
