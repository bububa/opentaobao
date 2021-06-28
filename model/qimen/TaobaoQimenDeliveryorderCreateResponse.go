package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发货单创建接口 APIResponse
taobao.qimen.deliveryorder.create

taobao.qimen.deliveryorder.create
*/
type TaobaoQimenDeliveryorderCreateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenDeliveryorderCreateResponse `json:"qimen_deliveryorder_create_response,omitempty"` 
    TaobaoQimenDeliveryorderCreateResponse
}

/* model for simplify = false
type TaobaoQimenDeliveryorderCreateResponse struct {

    // 
    
    Response  *struct {
        DeliveryOrderCreateResponse  *DeliveryOrderCreateResponse `json:"delivery_order_create_response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenDeliveryorderCreateResponse struct {

    // 
    Response   *DeliveryOrderCreateResponse `json:"response,omitempty"`

}
