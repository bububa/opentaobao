package ieagency

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新查询退票申请单列表 APIResponse
taobao.alitrip.ie.agent.refund.new.getlist

查询商家国际机票退票申请单列表
*/
type TaobaoAlitripIeAgentRefundNewGetlistAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripIeAgentRefundNewGetlistResponse
}

type TaobaoAlitripIeAgentRefundNewGetlistResponse struct {
    XMLName xml.Name `xml:"alitrip_ie_agent_refund_new_getlist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *RefundOrderQueryListRs `json:"result,omitempty" xml:"result,omitempty"`

    
}
