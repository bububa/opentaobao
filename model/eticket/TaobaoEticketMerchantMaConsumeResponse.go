package eticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证核销接口 APIResponse
taobao.eticket.merchant.ma.consume

电子凭证核销接口
*/
type TaobaoEticketMerchantMaConsumeAPIResponse struct {
    model.CommonResponse
    Response *TaobaoEticketMerchantMaConsumeResponse `json:"taobao_eticket_merchant_ma_consume_response,omitempty"`
}

type TaobaoEticketMerchantMaConsumeResponse struct {

    // 系统自动生成
    RespBody   *ConsumeMaCallbackResp `json:"resp_body,omitempty"`

    // 子结果码
    RetCode   string `json:"ret_code,omitempty"`

    // 子结果信息
    RetMsg   string `json:"ret_msg,omitempty"`

}
