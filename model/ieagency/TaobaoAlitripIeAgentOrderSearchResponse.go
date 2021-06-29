package ieagency

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票】订单列表查询 APIResponse
taobao.alitrip.ie.agent.order.search

根据指定条件查询国际机票订单列表
*/
type TaobaoAlitripIeAgentOrderSearchAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripIeAgentOrderSearchResponse
}

type TaobaoAlitripIeAgentOrderSearchResponse struct {
    XMLName xml.Name `xml:"alitrip_ie_agent_order_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 订单列表
    
    BaseOrderVos   []IeBaseOrderVo `json:"base_order_vos,omitempty" xml:"base_order_vos>ie_base_order_vo,omitempty"`
    
    
    // 请求成功标识
    
    QuerySuccess   bool `json:"query_success,omitempty" xml:"query_success,omitempty"`

    
    // 是否可以翻页查询
    
    HasNext   bool `json:"has_next,omitempty" xml:"has_next,omitempty"`

    
}
