package flight

import (
    "github.com/bububa/opentaobao/model"
)

/* 
廉航辅营正向订单查询详情接口 APIResponse
alitrip.tripvp.agent.order.get

【国际机票】查询辅营订单详情
*/
type AlitripTripvpAgentOrderGetAPIResponse struct {
    model.CommonResponse
    // Response *AlitripTripvpAgentOrderGetResponse `json:"alitrip_tripvp_agent_order_get_response,omitempty"` 
    AlitripTripvpAgentOrderGetResponse
}

/* model for simplify = false
type AlitripTripvpAgentOrderGetResponse struct {

    // orderVO
    
    OrderVo  *struct {
        VirProOrderVo  *VirProOrderVo `json:"vir_pro_order_vo,omitempty"`
    } `json:"order_vo,omitempty"`
    

    // pageSize
    
    PageSize   int64 `json:"page_size,omitempty"`
    

}
*/

type AlitripTripvpAgentOrderGetResponse struct {

    // orderVO
    OrderVo   *VirProOrderVo `json:"order_vo,omitempty"`

    // pageSize
    PageSize   int64 `json:"page_size,omitempty"`

}
