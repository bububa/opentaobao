package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发货单查询接口 APIResponse
taobao.qimen.deliveryorder.query

ERP调用奇门的发货单查询接口，查询发货单详情
*/
type TaobaoQimenDeliveryorderQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenDeliveryorderQueryResponse `json:"qimen_deliveryorder_query_response,omitempty"` 
    TaobaoQimenDeliveryorderQueryResponse
}

/* model for simplify = false
type TaobaoQimenDeliveryorderQueryResponse struct {

    // 
    
    Response  *struct {
        DeliveryOrderQueryResponse  *DeliveryOrderQueryResponse `json:"delivery_order_query_response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenDeliveryorderQueryResponse struct {

    // 
    Response   *DeliveryOrderQueryResponse `json:"response,omitempty"`

}
