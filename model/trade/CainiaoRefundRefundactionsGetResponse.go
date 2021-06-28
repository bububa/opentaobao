package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
判断该订单能执行的逆向操作集合列表 APIResponse
cainiao.refund.refundactions.get

判断该订单能执行的逆向操作集合列表
*/
type CainiaoRefundRefundactionsGetAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoRefundRefundactionsGetResponse `json:"cainiao_refund_refundactions_get_response,omitempty"` 
    CainiaoRefundRefundactionsGetResponse
}

/* model for simplify = false
type CainiaoRefundRefundactionsGetResponse struct {

    // 返回结果对象
    
    Result  *struct {
        CainiaoRefundRefundactionsGetBizResult  *CainiaoRefundRefundactionsGetBizResult `json:"cainiao_refund_refundactions_get_biz_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type CainiaoRefundRefundactionsGetResponse struct {

    // 返回结果对象
    Result   *CainiaoRefundRefundactionsGetBizResult `json:"result,omitempty"`

}
