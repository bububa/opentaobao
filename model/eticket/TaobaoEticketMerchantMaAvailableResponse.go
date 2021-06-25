package eticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证核销前校验接口 APIResponse
taobao.eticket.merchant.ma.available

商家验码之前的调用接口，用来判断是否可以进行核销操作
*/
type TaobaoEticketMerchantMaAvailableAPIResponse struct {
    model.CommonResponse
    Response *TaobaoEticketMerchantMaAvailableResponse `json:"taobao_eticket_merchant_ma_available_response,omitempty"`
}

type TaobaoEticketMerchantMaAvailableResponse struct {

    // 回复结果
    RespBody   *ConsumeMaCallbackResp `json:"resp_body,omitempty"`

    // 子结果码
    RetCode   string `json:"ret_code,omitempty"`

    // 子结果信息
    RetMsg   string `json:"ret_msg,omitempty"`

}
