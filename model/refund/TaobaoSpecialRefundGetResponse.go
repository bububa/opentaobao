package refund

import (
    "github.com/bububa/opentaobao/model"
)

/* 
特殊部分退纠纷单查询 APIResponse
taobao.special.refund.get

获取单笔特殊部分退的纠纷单查询
*/
type TaobaoSpecialRefundGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSpecialRefundGetResponse `json:"taobao_special_refund_get_response,omitempty"`
}

type TaobaoSpecialRefundGetResponse struct {

    // 退款详情
    Refund   *Refund `json:"refund,omitempty"`

}
