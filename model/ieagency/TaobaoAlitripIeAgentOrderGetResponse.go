package ieagency

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票】查询订单详情 APIResponse
taobao.alitrip.ie.agent.order.get

根据订单ID查询订单详情
*/
type TaobaoAlitripIeAgentOrderGetAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripIeAgentOrderGetResponse
}

type TaobaoAlitripIeAgentOrderGetResponse struct {
    XMLName xml.Name `xml:"alitrip_ie_agent_order_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回数据
    
    OrderVo   *IeOrderVo `json:"order_vo,omitempty" xml:"order_vo,omitempty"`

    
}
