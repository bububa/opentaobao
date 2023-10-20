package ieagency

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentRefundSearchAPIResponse 卖家查询退票申请 API返回值
// taobao.alitrip.ie.agent.refund.search
//
// 卖家查询退票申请
type TaobaoAlitripIeAgentRefundSearchAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripIeAgentRefundSearchAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentRefundSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripIeAgentRefundSearchAPIResponseModel).Reset()
}

// TaobaoAlitripIeAgentRefundSearchAPIResponseModel is 卖家查询退票申请 成功返回结果
type TaobaoAlitripIeAgentRefundSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ie_agent_refund_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *QueryRefundTicketsRs `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentRefundSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripIeAgentRefundSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripIeAgentRefundSearchAPIResponse)
	},
}

// GetTaobaoAlitripIeAgentRefundSearchAPIResponse 从 sync.Pool 获取 TaobaoAlitripIeAgentRefundSearchAPIResponse
func GetTaobaoAlitripIeAgentRefundSearchAPIResponse() *TaobaoAlitripIeAgentRefundSearchAPIResponse {
	return poolTaobaoAlitripIeAgentRefundSearchAPIResponse.Get().(*TaobaoAlitripIeAgentRefundSearchAPIResponse)
}

// ReleaseTaobaoAlitripIeAgentRefundSearchAPIResponse 将 TaobaoAlitripIeAgentRefundSearchAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripIeAgentRefundSearchAPIResponse(v *TaobaoAlitripIeAgentRefundSearchAPIResponse) {
	v.Reset()
	poolTaobaoAlitripIeAgentRefundSearchAPIResponse.Put(v)
}
