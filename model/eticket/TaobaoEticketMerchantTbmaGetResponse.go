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
    // Response *TaobaoEticketMerchantTbmaGetResponse `json:"eticket_merchant_tbma_get_response,omitempty"` 
    TaobaoEticketMerchantTbmaGetResponse
}

/* model for simplify = false
type TaobaoEticketMerchantTbmaGetResponse struct {

    // respBody
    
    RespBody  *struct {
        QueryTbMaCallbackResp  *QueryTbMaCallbackResp `json:"query_tb_ma_callback_resp,omitempty"`
    } `json:"resp_body,omitempty"`
    

    // subCode
    
    RetCode   string `json:"ret_code,omitempty"`
    

    // subMsg
    
    RetMsg   string `json:"ret_msg,omitempty"`
    

}
*/

type TaobaoEticketMerchantTbmaGetResponse struct {

    // respBody
    RespBody   *QueryTbMaCallbackResp `json:"resp_body,omitempty"`

    // subCode
    RetCode   string `json:"ret_code,omitempty"`

    // subMsg
    RetMsg   string `json:"ret_msg,omitempty"`

}
