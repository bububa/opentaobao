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
    Response *TaobaoAlitripCarOrderQueryResponse `json:"taobao_alitrip_car_order_query_response,omitempty"`
}

type TaobaoAlitripCarOrderQueryResponse struct {

    // 订单结果
    FirstResult   *OrderQueryRsp `json:"first_result,omitempty"`

}
