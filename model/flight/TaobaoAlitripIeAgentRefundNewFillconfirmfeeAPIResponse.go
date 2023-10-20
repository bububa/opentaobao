package flight

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponse 新模型-回填申请单费用 API返回值
// taobao.alitrip.ie.agent.refund.new.fillconfirmfee
//
// 1. 回填退票费用
type TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponseModel).Reset()
}

// TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponseModel is 新模型-回填申请单费用 成功返回结果
type TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ie_agent_refund_new_fillconfirmfee_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *RefundOrderFillConfirmFeeRs `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponse)
	},
}

// GetTaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponse 从 sync.Pool 获取 TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponse
func GetTaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponse() *TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponse {
	return poolTaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponse.Get().(*TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponse)
}

// ReleaseTaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponse 将 TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponse(v *TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponse) {
	v.Reset()
	poolTaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponse.Put(v)
}
