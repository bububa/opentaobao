package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
退货退款操作的展示信息(展现给买家) APIResponse
cainiao.refund.refundactions.display

退货退款操作的展示信息(展现给买家)
*/
type CainiaoRefundRefundactionsDisplayAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoRefundRefundactionsDisplayResponse `json:"cainiao_refund_refundactions_display_response,omitempty"` 
    CainiaoRefundRefundactionsDisplayResponse
}

/* model for simplify = false
type CainiaoRefundRefundactionsDisplayResponse struct {

    // 返回结果对象
    
    Result  *struct {
        CainiaoRefundRefundactionsDisplayBizResult  *CainiaoRefundRefundactionsDisplayBizResult `json:"cainiao_refund_refundactions_display_biz_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type CainiaoRefundRefundactionsDisplayResponse struct {

    // 返回结果对象
    Result   *CainiaoRefundRefundactionsDisplayBizResult `json:"result,omitempty"`

}
