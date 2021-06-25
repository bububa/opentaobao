package car

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家订单状态改变通知接口（神州专车接口） APIResponse
taobao.alitrip.car.order.status

商家订单状态改变通知接口，神州专车专用接口！
*/
type TaobaoAlitripCarOrderStatusAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAlitripCarOrderStatusResponse `json:"taobao_alitrip_car_order_status_response,omitempty"`
}

type TaobaoAlitripCarOrderStatusResponse struct {

    // 根据站点名称查询产品
    Result   *TaobaoAlitripCarOrderStatusApiResult `json:"result,omitempty"`

}
