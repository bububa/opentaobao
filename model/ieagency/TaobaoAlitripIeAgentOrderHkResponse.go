package ieagency

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票】手工预定回填PNR APIResponse
taobao.alitrip.ie.agent.order.hk

代理商通过手工预定PNR，并回填。
*/
type TaobaoAlitripIeAgentOrderHkAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripIeAgentOrderHkResponse
}

type TaobaoAlitripIeAgentOrderHkResponse struct {
    XMLName xml.Name `xml:"alitrip_ie_agent_order_hk_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否回填成功true：成功 false:失败
    
    HkSuccess   bool `json:"hk_success,omitempty" xml:"hk_success,omitempty"`

    
}
