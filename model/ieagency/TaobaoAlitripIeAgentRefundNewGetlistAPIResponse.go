package ieagency

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentRefundNewGetlistAPIResponse 新查询退票申请单列表 API返回值
// taobao.alitrip.ie.agent.refund.new.getlist
//
// 查询商家国际机票退票申请单列表
type TaobaoAlitripIeAgentRefundNewGetlistAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripIeAgentRefundNewGetlistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentRefundNewGetlistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripIeAgentRefundNewGetlistAPIResponseModel).Reset()
}

// TaobaoAlitripIeAgentRefundNewGetlistAPIResponseModel is 新查询退票申请单列表 成功返回结果
type TaobaoAlitripIeAgentRefundNewGetlistAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ie_agent_refund_new_getlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *RefundOrderQueryListRs `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentRefundNewGetlistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripIeAgentRefundNewGetlistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripIeAgentRefundNewGetlistAPIResponse)
	},
}

// GetTaobaoAlitripIeAgentRefundNewGetlistAPIResponse 从 sync.Pool 获取 TaobaoAlitripIeAgentRefundNewGetlistAPIResponse
func GetTaobaoAlitripIeAgentRefundNewGetlistAPIResponse() *TaobaoAlitripIeAgentRefundNewGetlistAPIResponse {
	return poolTaobaoAlitripIeAgentRefundNewGetlistAPIResponse.Get().(*TaobaoAlitripIeAgentRefundNewGetlistAPIResponse)
}

// ReleaseTaobaoAlitripIeAgentRefundNewGetlistAPIResponse 将 TaobaoAlitripIeAgentRefundNewGetlistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripIeAgentRefundNewGetlistAPIResponse(v *TaobaoAlitripIeAgentRefundNewGetlistAPIResponse) {
	v.Reset()
	poolTaobaoAlitripIeAgentRefundNewGetlistAPIResponse.Put(v)
}
