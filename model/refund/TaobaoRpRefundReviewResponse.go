package refund

import (
    "github.com/bububa/opentaobao/model"
)

/* 
审核退款单 APIResponse
taobao.rp.refund.review

审核退款单，标志是否可用于批量退款，目前仅支持天猫订单。
*/
type TaobaoRpRefundReviewAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoRpRefundReviewResponse `json:"rp_refund_review_response,omitempty"` 
    TaobaoRpRefundReviewResponse
}

/* model for simplify = false
type TaobaoRpRefundReviewResponse struct {

    // success
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoRpRefundReviewResponse struct {

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

}
