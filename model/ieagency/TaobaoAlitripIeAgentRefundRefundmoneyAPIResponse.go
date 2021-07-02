package ieagency

import (
	"encoding/xml"

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

// TaobaoAlitripIeAgentRefundRefundmoneyAPIResponseModel is 确认退款 成功返回结果
type TaobaoAlitripIeAgentRefundRefundmoneyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ie_agent_refund_refundmoney_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *RefundMoneyNoPasswordRs `json:"result,omitempty" xml:"result,omitempty"`
}
