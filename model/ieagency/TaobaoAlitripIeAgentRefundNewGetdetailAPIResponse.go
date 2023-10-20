package ieagency

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentRefundNewGetdetailAPIResponse 查询申请单详情(新版) API返回值
// taobao.alitrip.ie.agent.refund.new.getdetail
//
// 查询申请单详情
type TaobaoAlitripIeAgentRefundNewGetdetailAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripIeAgentRefundNewGetdetailAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentRefundNewGetdetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripIeAgentRefundNewGetdetailAPIResponseModel).Reset()
}

// TaobaoAlitripIeAgentRefundNewGetdetailAPIResponseModel is 查询申请单详情(新版) 成功返回结果
type TaobaoAlitripIeAgentRefundNewGetdetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ie_agent_refund_new_getdetail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *RefundOrderQueryDetailRs `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentRefundNewGetdetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripIeAgentRefundNewGetdetailAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripIeAgentRefundNewGetdetailAPIResponse)
	},
}

// GetTaobaoAlitripIeAgentRefundNewGetdetailAPIResponse 从 sync.Pool 获取 TaobaoAlitripIeAgentRefundNewGetdetailAPIResponse
func GetTaobaoAlitripIeAgentRefundNewGetdetailAPIResponse() *TaobaoAlitripIeAgentRefundNewGetdetailAPIResponse {
	return poolTaobaoAlitripIeAgentRefundNewGetdetailAPIResponse.Get().(*TaobaoAlitripIeAgentRefundNewGetdetailAPIResponse)
}

// ReleaseTaobaoAlitripIeAgentRefundNewGetdetailAPIResponse 将 TaobaoAlitripIeAgentRefundNewGetdetailAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripIeAgentRefundNewGetdetailAPIResponse(v *TaobaoAlitripIeAgentRefundNewGetdetailAPIResponse) {
	v.Reset()
	poolTaobaoAlitripIeAgentRefundNewGetdetailAPIResponse.Put(v)
}
