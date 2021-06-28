package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询采购单退款信息 APIResponse
taobao.fenxiao.refund.get

分销商或供应商可以查询某子单的退款信息，以及下游订单的退款信息
*/
type TaobaoFenxiaoRefundGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoFenxiaoRefundGetResponse `json:"fenxiao_refund_get_response,omitempty"` 
    TaobaoFenxiaoRefundGetResponse
}

/* model for simplify = false
type TaobaoFenxiaoRefundGetResponse struct {

    // 退款详情
    
    RefundDetail  *struct {
        RefundDetail  *RefundDetail `json:"refund_detail,omitempty"`
    } `json:"refund_detail,omitempty"`
    

}
*/

type TaobaoFenxiaoRefundGetResponse struct {

    // 退款详情
    RefundDetail   *RefundDetail `json:"refund_detail,omitempty"`

}
