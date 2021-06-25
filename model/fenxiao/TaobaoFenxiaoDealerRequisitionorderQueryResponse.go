package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
按编号查询采购申请/经销采购单 APIResponse
taobao.fenxiao.dealer.requisitionorder.query

按编号查询采购申请/经销采购单，目前支持供应商和分销商查询。
*/
type TaobaoFenxiaoDealerRequisitionorderQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoFenxiaoDealerRequisitionorderQueryResponse `json:"taobao_fenxiao_dealer_requisitionorder_query_response,omitempty"`
}

type TaobaoFenxiaoDealerRequisitionorderQueryResponse struct {

    // 经销采购单结果列表
    DealerOrders   []DealerOrder `json:"dealer_orders,omitempty"`

}
