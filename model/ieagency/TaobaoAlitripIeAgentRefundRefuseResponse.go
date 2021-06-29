package ieagency

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
拒绝退票申请 APIResponse
taobao.alitrip.ie.agent.refund.refuse

卖家拒绝退票退票申请
*/
type TaobaoAlitripIeAgentRefundRefuseAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripIeAgentRefundRefuseResponse
}

type TaobaoAlitripIeAgentRefundRefuseResponse struct {
    XMLName xml.Name `xml:"alitrip_ie_agent_refund_refuse_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *RefuseRefundTicketRs `json:"result,omitempty" xml:"result,omitempty"`

    
}
