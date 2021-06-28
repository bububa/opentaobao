package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
确认收款 APIResponse
taobao.fenxiao.order.confirm.paid

供应商确认收款（非支付宝交易）。
*/
type TaobaoFenxiaoOrderConfirmPaidAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoFenxiaoOrderConfirmPaidResponse `json:"fenxiao_order_confirm_paid_response,omitempty"` 
    TaobaoFenxiaoOrderConfirmPaidResponse
}

/* model for simplify = false
type TaobaoFenxiaoOrderConfirmPaidResponse struct {

    // 确认结果成功与否
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoFenxiaoOrderConfirmPaidResponse struct {

    // 确认结果成功与否
    IsSuccess   bool `json:"is_success,omitempty"`

}
