package ieagency

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
确认退款 APIResponse
taobao.alitrip.ie.agent.refund.refundmoney

卖家进行退款操作
*/
type TaobaoAlitripIeAgentRefundRefundmoneyAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripIeAgentRefundRefundmoneyResponse
}

type TaobaoAlitripIeAgentRefundRefundmoneyResponse struct {
    XMLName xml.Name `xml:"alitrip_ie_agent_refund_refundmoney_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *RefundMoneyNoPasswordRs `json:"result,omitempty" xml:"result,omitempty"`

    
}
