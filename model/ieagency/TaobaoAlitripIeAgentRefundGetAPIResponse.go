package ieagency

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentRefundGetAPIResponse 获取退票申请详情 API返回值
// taobao.alitrip.ie.agent.refund.get
//
// 获取退票申请详情
type TaobaoAlitripIeAgentRefundGetAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripIeAgentRefundGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentRefundGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripIeAgentRefundGetAPIResponseModel).Reset()
}

// TaobaoAlitripIeAgentRefundGetAPIResponseModel is 获取退票申请详情 成功返回结果
type TaobaoAlitripIeAgentRefundGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ie_agent_refund_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *QueryRefundTicketDetailRs `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentRefundGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripIeAgentRefundGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripIeAgentRefundGetAPIResponse)
	},
}

// GetTaobaoAlitripIeAgentRefundGetAPIResponse 从 sync.Pool 获取 TaobaoAlitripIeAgentRefundGetAPIResponse
func GetTaobaoAlitripIeAgentRefundGetAPIResponse() *TaobaoAlitripIeAgentRefundGetAPIResponse {
	return poolTaobaoAlitripIeAgentRefundGetAPIResponse.Get().(*TaobaoAlitripIeAgentRefundGetAPIResponse)
}

// ReleaseTaobaoAlitripIeAgentRefundGetAPIResponse 将 TaobaoAlitripIeAgentRefundGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripIeAgentRefundGetAPIResponse(v *TaobaoAlitripIeAgentRefundGetAPIResponse) {
	v.Reset()
	poolTaobaoAlitripIeAgentRefundGetAPIResponse.Put(v)
}
