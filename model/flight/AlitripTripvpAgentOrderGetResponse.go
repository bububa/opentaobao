package flight

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
廉航辅营正向订单查询详情接口 APIResponse
alitrip.tripvp.agent.order.get

【国际机票】查询辅营订单详情
*/
type AlitripTripvpAgentOrderGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_tripvp_agent_order_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // orderVO
    
    OrderVo   *VirProOrderVo `json:"order_vo,omitempty" xml:"