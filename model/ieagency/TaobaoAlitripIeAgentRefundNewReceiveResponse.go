package ieagency

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家退票受理申请(对外) APIResponse
taobao.alitrip.ie.agent.refund.new.receive

允许代理商通过top接口受理退票申请
*/
type TaobaoAlitripIeAgentRefundNewReceiveAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripIeAgentRefundNewReceiveResponse
}

type TaobaoAlitripIeAgentRefundNewReceiveResponse struct {
    XMLName xml.Name `xml:"alitrip_ie_agent_refund_new_receive_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ReceiveRefundTicketRs `json:"result,omitempty" xml:"result,omitempty"`

    
}
