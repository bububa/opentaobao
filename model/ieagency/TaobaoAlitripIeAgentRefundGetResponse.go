package ieagency

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取退票申请详情 APIResponse
taobao.alitrip.ie.agent.refund.get

获取退票申请详情
*/
type TaobaoAlitripIeAgentRefundGetAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripIeAgentRefundGetResponse
}

type TaobaoAlitripIeAgentRefundGetResponse struct {
    XMLName xml.Name `xml:"alitrip_ie_agent_refund_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *QueryRefundTicketDetailRs `json:"result,omitempty" xml:"result,omitempty"`

    
}
