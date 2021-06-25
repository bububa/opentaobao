package eticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证冲正接口 APIResponse
taobao.eticket.merchant.ma.reverse

电子凭证平台冲正接口
*/
type TaobaoEticketMerchantMaReverseAPIResponse struct {
    model.CommonResponse
    Response *TaobaoEticketMerchantMaReverseResponse `json:"taobao_eticket_merchant_ma_reverse_response,omitempty"`
}

type TaobaoEticketMerchantMaReverseResponse struct {

    // 回复结果
    RespBody   *ReverseMaCallbackResp `json:"resp_body,omitempty"`

    // 子结果码
    RetCode   string `json:"ret_code,omitempty"`

    // 子结果信息
    RetMsg   string `json:"ret_msg,omitempty"`

}
