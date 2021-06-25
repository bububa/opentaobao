package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量查询采购退款 APIResponse
taobao.fenxiao.refund.query

供应商按查询条件批量查询代销采购退款
*/
type TaobaoFenxiaoRefundQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoFenxiaoRefundQueryResponse `json:"taobao_fenxiao_refund_query_response,omitempty"`
}

type TaobaoFenxiaoRefundQueryResponse struct {

    // 按查询条件查到的记录总数
    TotalResults   int64 `json:"total_results,omitempty"`

    // 代销采购退款列表
    RefundList   []RefundDetail `json:"refund_list,omitempty"`

}
