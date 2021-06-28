package refund

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取单笔退款详情 APIResponse
taobao.refund.get

获取单笔退款详情
*/
type TaobaoRefundGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoRefundGetResponse `json:"refund_get_response,omitempty"` 
    TaobaoRefundGetResponse
}

/* model for simplify = false
type TaobaoRefundGetResponse struct {

    // 退款详情
    
    Refund  *struct {
        Refund  *Refund `json:"refund,omitempty"`
    } `json:"refund,omitempty"`
    

}
*/

type TaobaoRefundGetResponse struct {

    // 退款详情
    Refund   *Refund `json:"refund,omitempty"`

}
