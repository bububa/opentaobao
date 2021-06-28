package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
无需物流（虚拟）发货处理 APIResponse
taobao.logistics.dummy.send

用户调用该接口可实现无需物流（虚拟）发货,使用该接口发货，交易订单状态会直接变成卖家已发货
*/
type TaobaoLogisticsDummySendAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoLogisticsDummySendResponse `json:"logistics_dummy_send_response,omitempty"` 
    TaobaoLogisticsDummySendResponse
}

/* model for simplify = false
type TaobaoLogisticsDummySendResponse struct {

    // 返回发货是否成功is_success
    
    Shipping  *struct {
        Shipping  *Shipping `json:"shipping,omitempty"`
    } `json:"shipping,omitempty"`
    

}
*/

type TaobaoLogisticsDummySendResponse struct {

    // 返回发货是否成功is_success
    Shipping   *Shipping `json:"shipping,omitempty"`

}
