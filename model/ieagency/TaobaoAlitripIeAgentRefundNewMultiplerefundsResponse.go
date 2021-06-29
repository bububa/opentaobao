package ieagency

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
补退接口 APIResponse
taobao.alitrip.ie.agent.refund.new.multiplerefunds

1. 补退接口， 可以进行多次退款
*/
type TaobaoAlitripIeAgentRefundNewMultiplerefundsAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripIeAgentRefundNewMultiplerefundsResponse
}

type TaobaoAlitripIeAgentRefundNewMultiplerefundsResponse struct {
    XMLName xml.Name `xml:"alitrip_ie_agent_refund_new_multiplerefunds_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *RefundOrderMultipleRefundsRs `json:"result,omitempty" xml:"result,omitempty"`

    
}
