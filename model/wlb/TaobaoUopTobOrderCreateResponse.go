package wlb

import (
    "github.com/bububa/opentaobao/model"
)

/* 
ToB仓储发货 APIResponse
taobao.uop.tob.order.create

ToB仓储发货
*/
type TaobaoUopTobOrderCreateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoUopTobOrderCreateResponse `json:"uop_tob_order_create_response,omitempty"` 
    TaobaoUopTobOrderCreateResponse
}

/* model for simplify = false
type TaobaoUopTobOrderCreateResponse struct {

    // flag
    
    Flag   string `json:"flag,omitempty"`
    

    // message
    
    Message   string `json:"message,omitempty"`
    

    // 订单
    
    DeliveryOrders  struct {
        Deliveryorder  []Deliveryorder `json:"deliveryorder,omitempty"`
    } `json:"delivery_orders,omitempty"`
    

}
*/

type TaobaoUopTobOrderCreateResponse struct {

    // flag
    Flag   string `json:"flag,omitempty"`

    // message
    Message   string `json:"message,omitempty"`

    // 订单
    DeliveryOrders   []Deliveryorder `json:"delivery_orders,omitempty"`

}
