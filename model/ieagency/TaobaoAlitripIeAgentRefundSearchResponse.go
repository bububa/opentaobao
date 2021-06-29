package ieagency

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家查询退票申请 APIResponse
taobao.alitrip.ie.agent.refund.search

卖家查询退票申请
*/
type TaobaoAlitripIeAgentRefundSearchAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripIeAgentRefundSearchResponse
}

type TaobaoAlitripIeAgentRefundSearchResponse struct {
    XMLName xml.Name `xml:"alitrip_ie_agent_refund_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *QueryRefundTicketsRs `json:"result,omitempty" xml:"result,omitempty"`

    
}
