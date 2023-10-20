package ieagency

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentRefundNewReceiveAPIResponse 商家退票受理申请(对外) API返回值
// taobao.alitrip.ie.agent.refund.new.receive
//
// 允许代理商通过top接口受理退票申请
type TaobaoAlitripIeAgentRefundNewReceiveAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripIeAgentRefundNewReceiveAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentRefundNewReceiveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripIeAgentRefundNewReceiveAPIResponseModel).Reset()
}

// TaobaoAlitripIeAgentRefundNewReceiveAPIResponseModel is 商家退票受理申请(对外) 成功返回结果
type TaobaoAlitripIeAgentRefundNewReceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ie_agent_refund_new_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ReceiveRefundTicketRs `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentRefundNewReceiveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripIeAgentRefundNewReceiveAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripIeAgentRefundNewReceiveAPIResponse)
	},
}

// GetTaobaoAlitripIeAgentRefundNewReceiveAPIResponse 从 sync.Pool 获取 TaobaoAlitripIeAgentRefundNewReceiveAPIResponse
func GetTaobaoAlitripIeAgentRefundNewReceiveAPIResponse() *TaobaoAlitripIeAgentRefundNewReceiveAPIResponse {
	return poolTaobaoAlitripIeAgentRefundNewReceiveAPIResponse.Get().(*TaobaoAlitripIeAgentRefundNewReceiveAPIResponse)
}

// ReleaseTaobaoAlitripIeAgentRefundNewReceiveAPIResponse 将 TaobaoAlitripIeAgentRefundNewReceiveAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripIeAgentRefundNewReceiveAPIResponse(v *TaobaoAlitripIeAgentRefundNewReceiveAPIResponse) {
	v.Reset()
	poolTaobaoAlitripIeAgentRefundNewReceiveAPIResponse.Put(v)
}
