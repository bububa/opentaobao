package ieagency

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentRefundRefuseAPIResponse 拒绝退票申请 API返回值
// taobao.alitrip.ie.agent.refund.refuse
//
// 卖家拒绝退票退票申请
type TaobaoAlitripIeAgentRefundRefuseAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripIeAgentRefundRefuseAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentRefundRefuseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripIeAgentRefundRefuseAPIResponseModel).Reset()
}

// TaobaoAlitripIeAgentRefundRefuseAPIResponseModel is 拒绝退票申请 成功返回结果
type TaobaoAlitripIeAgentRefundRefuseAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ie_agent_refund_refuse_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *RefuseRefundTicketRs `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentRefundRefuseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripIeAgentRefundRefuseAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripIeAgentRefundRefuseAPIResponse)
	},
}

// GetTaobaoAlitripIeAgentRefundRefuseAPIResponse 从 sync.Pool 获取 TaobaoAlitripIeAgentRefundRefuseAPIResponse
func GetTaobaoAlitripIeAgentRefundRefuseAPIResponse() *TaobaoAlitripIeAgentRefundRefuseAPIResponse {
	return poolTaobaoAlitripIeAgentRefundRefuseAPIResponse.Get().(*TaobaoAlitripIeAgentRefundRefuseAPIResponse)
}

// ReleaseTaobaoAlitripIeAgentRefundRefuseAPIResponse 将 TaobaoAlitripIeAgentRefundRefuseAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripIeAgentRefundRefuseAPIResponse(v *TaobaoAlitripIeAgentRefundRefuseAPIResponse) {
	v.Reset()
	poolTaobaoAlitripIeAgentRefundRefuseAPIResponse.Put(v)
}
