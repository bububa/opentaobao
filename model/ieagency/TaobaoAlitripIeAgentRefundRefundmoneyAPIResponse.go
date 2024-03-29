package ieagency

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentRefundRefundmoneyAPIResponse 确认退款 API返回值
// taobao.alitrip.ie.agent.refund.refundmoney
//
// 卖家进行退款操作
type TaobaoAlitripIeAgentRefundRefundmoneyAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripIeAgentRefundRefundmoneyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentRefundRefundmoneyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripIeAgentRefundRefundmoneyAPIResponseModel).Reset()
}

// TaobaoAlitripIeAgentRefundRefundmoneyAPIResponseModel is 确认退款 成功返回结果
type TaobaoAlitripIeAgentRefundRefundmoneyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ie_agent_refund_refundmoney_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *RefundMoneyNoPasswordRs `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentRefundRefundmoneyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripIeAgentRefundRefundmoneyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripIeAgentRefundRefundmoneyAPIResponse)
	},
}

// GetTaobaoAlitripIeAgentRefundRefundmoneyAPIResponse 从 sync.Pool 获取 TaobaoAlitripIeAgentRefundRefundmoneyAPIResponse
func GetTaobaoAlitripIeAgentRefundRefundmoneyAPIResponse() *TaobaoAlitripIeAgentRefundRefundmoneyAPIResponse {
	return poolTaobaoAlitripIeAgentRefundRefundmoneyAPIResponse.Get().(*TaobaoAlitripIeAgentRefundRefundmoneyAPIResponse)
}

// ReleaseTaobaoAlitripIeAgentRefundRefundmoneyAPIResponse 将 TaobaoAlitripIeAgentRefundRefundmoneyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripIeAgentRefundRefundmoneyAPIResponse(v *TaobaoAlitripIeAgentRefundRefundmoneyAPIResponse) {
	v.Reset()
	poolTaobaoAlitripIeAgentRefundRefundmoneyAPIResponse.Put(v)
}
