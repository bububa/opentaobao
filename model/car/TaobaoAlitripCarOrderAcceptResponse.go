package car

import (
    "github.com/bububa/opentaobao/model"
)

/* 
确认接单 APIResponse
taobao.alitrip.car.order.accept

用来接收服务商确认接单信息
*/
type TaobaoAlitripCarOrderAcceptAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAlitripCarOrderAcceptResponse `json:"alitrip_car_order_accept_response,omitempty"` 
    TaobaoAlitripCarOrderAcceptResponse
}

/* model for simplify = false
type TaobaoAlitripCarOrderAcceptResponse struct {

    // 根据站点名称查询产品
    
    Result  *struct {
        TaobaoAlitripCarOrderAcceptApiResult  *TaobaoAlitripCarOrderAcceptApiResult `json:"taobao_alitrip_car_order_accept_api_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoAlitripCarOrderAcceptResponse struct {

    // 根据站点名称查询产品
    Result   *TaobaoAlitripCarOrderAcceptApiResult `json:"result,omitempty"`

}
