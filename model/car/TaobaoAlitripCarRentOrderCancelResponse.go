package car

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
租车-取消订单 APIResponse
taobao.alitrip.car.rent.order.cancel

服务商主动取消用户订单或者拒绝取消订单.
*/
type TaobaoAlitripCarRentOrderCancelAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_car_rent_order_cancel_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误消息
    
    Message   string `json:"message,omitempty" xml:"