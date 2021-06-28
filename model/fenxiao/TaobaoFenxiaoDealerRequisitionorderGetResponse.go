package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量查询采购申请/经销采购单 APIResponse
taobao.fenxiao.dealer.requisitionorder.get

批量查询采购申请/经销采购单，目前支持供应商和分销商查询
*/
type TaobaoFenxiaoDealerRequisitionorderGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoFenxiaoDealerRequisitionorderGetResponse `json:"fenxiao_dealer_requisitionorder_get_response,omitempty"` 
    TaobaoFenxiaoDealerRequisitionorderGetResponse
}

/* model for simplify = false
type TaobaoFenxiaoDealerRequisitionorderGetResponse struct {

    // 按查询条件查到的记录总数
    
    TotalResults   int64 `json:"total_results,omitempty"`
    

    // 采购申请/经销采购单结果列表
    
    DealerOrders  struct {
        DealerOrder  []DealerOrder `json:"dealer_order,omitempty"`
    } `json:"dealer_orders,omitempty"`
    

}
*/

type TaobaoFenxiaoDealerRequisitionorderGetResponse struct {

    // 按查询条件查到的记录总数
    TotalResults   int64 `json:"total_results,omitempty"`

    // 采购申请/经销采购单结果列表
    DealerOrders   []DealerOrder `json:"dealer_orders,omitempty"`

}
