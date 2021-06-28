package car

import (
    "github.com/bububa/opentaobao/model"
)

/* 
飞猪订单状态查询接口 APIResponse
taobao.alitrip.car.order.query

提供给直连商家查询在飞猪平台上产生的订单
*/
type TaobaoAlitripCarOrderQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAlitripCarOrderQueryResponse `json:"alitrip_car_order_query_response,omitempty"` 
    TaobaoAlitripCarOrderQueryResponse
}

/* model for simplify = false
type TaobaoAlitripCarOrderQueryResponse struct {

    // 订单结果
    
    FirstResult  *struct {
        OrderQueryRsp  *OrderQueryRsp `json:"order_query_rsp,omitempty"`
    } `json:"first_result,omitempty"`
    

}
*/

type TaobaoAlitripCarOrderQueryResponse struct {

    // 订单结果
    FirstResult   *OrderQueryRsp `json:"first_result,omitempty"`

}
