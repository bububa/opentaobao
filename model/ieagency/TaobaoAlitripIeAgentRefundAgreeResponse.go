package ieagency

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
同意退票 APIResponse
taobao.alitrip.ie.agent.refund.agree

卖家同意买家退票申请
*/
type TaobaoAlitripIeAgentRefundAgreeAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripIeAgentRefundAgreeResponse
}

type TaobaoAlitripIeAgentRefundAgreeResponse struct {
    XMLName xml.Name `xml:"alitrip_ie_agent_refund_agree_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AgreeRefundTicketRs `json:"result,omitempty" xml:"result,omitempty"`

    
}
