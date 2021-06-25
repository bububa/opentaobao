package eticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
外部合作商家重发电子凭证回调接口 APIResponse
taobao.vmarket.eticket.resend

外部合作商家重发电子凭证回调接口
*/
type TaobaoVmarketEticketResendAPIResponse struct {
    model.CommonResponse
    Response *TaobaoVmarketEticketResendResponse `json:"taobao_vmarket_eticket_resend_response,omitempty"`
}

type TaobaoVmarketEticketResendResponse struct {

    // 0:失败，1:成功
    RetCode   int64 `json:"ret_code,omitempty"`

}
