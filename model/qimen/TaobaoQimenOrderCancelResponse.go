package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
单据取消接口 APIResponse
taobao.qimen.order.cancel

ERP调用奇门的接口,取消创建单据操作。场景介绍：ERP主动发起取消某些创建的单据。如入库单、出库单、退货单等；所有的场景
*/
type TaobaoQimenOrderCancelAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenOrderCancelResponse `json:"qimen_order_cancel_response,omitempty"` 
    TaobaoQimenOrderCancelResponse
}

/* model for simplify = false
type TaobaoQimenOrderCancelResponse struct {

    // 
    
    Response  *struct {
        OrderCancelResponse  *OrderCancelResponse `json:"order_cancel_response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenOrderCancelResponse struct {

    // 
    Response   *OrderCancelResponse `json:"response,omitempty"`

}
