package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
单据挂起（恢复）接口 APIResponse
taobao.qimen.order.pending

ERP调用奇门的接口,挂起某些创建的单据;场景介绍：ERP主动发起挂起（恢复）某些创建的单据，如入库单、出库单、退货单等
*/
type TaobaoQimenOrderPendingAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenOrderPendingResponse `json:"taobao_qimen_order_pending_response,omitempty"`
}

type TaobaoQimenOrderPendingResponse struct {

    // 
    Response   *OrderPendingResponse `json:"response,omitempty"`

}
