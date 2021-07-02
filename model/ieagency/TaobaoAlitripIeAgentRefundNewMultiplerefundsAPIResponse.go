package ieagency

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentRefundNewMultiplerefundsAPIResponse 补退接口 API返回值
// taobao.alitrip.ie.agent.refund.new.multiplerefunds
//
// 1. 补退接口， 可以进行多次退款
type TaobaoAlitripIeAgentRefundNewMultiplerefundsAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripIeAgentRefundNewMultiplerefundsAPIResponseModel
}

// TaobaoAlitripIeAgentRefundNewMultiplerefundsAPIResponseModel is 补退接口 成功返回结果
type TaobaoAlitripIeAgentRefundNewMultiplerefundsAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ie_agent_refund_new_multiplerefunds_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *RefundOrderMultipleRefundsRs `json:"result,omitempty" xml:"result,omitempty"`
}
