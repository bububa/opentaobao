package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
判断当前用户是否能对订单执行一些逆向操作，比如退货操作 APIResponse
cainiao.refund.refundactions.judgement

判断当前用户是否能对订单执行一些逆向操作，比如退货操作
*/
type CainiaoRefundRefundactionsJudgementAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoRefundRefundactionsJudgementResponse `json:"cainiao_refund_refundactions_judgement_response,omitempty"` 
    CainiaoRefundRefundactionsJudgementResponse
}

/* model for simplify = false
type CainiaoRefundRefundactionsJudgementResponse struct {

    // 返回结果对象
    
    Result  *struct {
        CainiaoRefundRefundactionsJudgementBizResult  *CainiaoRefundRefundactionsJudgementBizResult `json:"cainiao_refund_refundactions_judgement_biz_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type CainiaoRefundRefundactionsJudgementResponse struct {

    // 返回结果对象
    Result   *CainiaoRefundRefundactionsJudgementBizResult `json:"result,omitempty"`

}
