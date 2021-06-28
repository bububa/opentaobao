package refund

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
审核退款单 APIResponse
taobao.rp.refund.review

审核退款单，标志是否可用于批量退款，目前仅支持天猫订单。
*/
type TaobaoRpRefundReviewAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"rp_refund_review_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"