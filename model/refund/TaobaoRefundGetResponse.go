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
    Response *TaobaoRefundGetResponse `json:"taobao_refund_get_response,omitempty"`
}

type TaobaoRefundGetResponse struct {

    // 退款详情
    Refund   *Refund `json:"refund,omitempty"`

}
