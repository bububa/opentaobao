package train

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentGetRefundAPIResponse 代理商获取订单退票信息 API返回值
// taobao.train.agent.get.refund
//
// 代理商获取订单信息回调API
type TaobaoTrainAgentGetRefundAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentGetRefundAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTrainAgentGetRefundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTrainAgentGetRefundAPIResponseModel).Reset()
}

// TaobaoTrainAgentGetRefundAPIResponseModel is 代理商获取订单退票信息 成功返回结果
type TaobaoTrainAgentGetRefundAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_get_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	TopRefundApplyList string `json:"top_refund_apply_list,omitempty" xml:"top_refund_apply_list,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTrainAgentGetRefundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TopRefundApplyList = ""
}

var poolTaobaoTrainAgentGetRefundAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTrainAgentGetRefundAPIResponse)
	},
}

// GetTaobaoTrainAgentGetRefundAPIResponse 从 sync.Pool 获取 TaobaoTrainAgentGetRefundAPIResponse
func GetTaobaoTrainAgentGetRefundAPIResponse() *TaobaoTrainAgentGetRefundAPIResponse {
	return poolTaobaoTrainAgentGetRefundAPIResponse.Get().(*TaobaoTrainAgentGetRefundAPIResponse)
}

// ReleaseTaobaoTrainAgentGetRefundAPIResponse 将 TaobaoTrainAgentGetRefundAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTrainAgentGetRefundAPIResponse(v *TaobaoTrainAgentGetRefundAPIResponse) {
	v.Reset()
	poolTaobaoTrainAgentGetRefundAPIResponse.Put(v)
}
