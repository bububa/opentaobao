package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建经销采购申请 APIResponse
taobao.fenxiao.dealer.requisitionorder.create

创建经销采购申请
*/
type TaobaoFenxiaoDealerRequisitionorderCreateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoFenxiaoDealerRequisitionorderCreateResponse `json:"taobao_fenxiao_dealer_requisitionorder_create_response,omitempty"`
}

type TaobaoFenxiaoDealerRequisitionorderCreateResponse struct {

    // 经销采购申请编号
    DealerOrderId   int64 `json:"dealer_order_id,omitempty"`

}
