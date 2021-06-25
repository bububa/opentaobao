package eticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
码商查询淘宝码接口 APIResponse
taobao.eticket.merchant.tbma.get

码商查询淘宝码接口
*/
type TaobaoEticketMerchantTbmaGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoEticketMerchantTbmaGetResponse `json:"taobao_eticket_merchant_tbma_get_response,omitempty"`
}

type TaobaoEticketMerchantTbmaGetResponse struct {

    // respBody
    RespBody   *QueryTbMaCallbackResp `json:"resp_body,omitempty"`

    // subCode
    RetCode   string `json:"ret_code,omitempty"`

    // subMsg
    RetMsg   string `json:"ret_msg,omitempty"`

}
