package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
POS收银成功后订单同步 APIResponse
alibaba.mj.oc.pay

此API用于在银泰商场中，消费者在收银台收银/退款时， POS系统在收银或退款成功后，调用此接口进行订单同步
*/
type AlibabaMjOcPayAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaMjOcPayResponse `json:"alibaba_mj_oc_pay_response,omitempty"` 
    AlibabaMjOcPayResponse
}

/* model for simplify = false
type AlibabaMjOcPayResponse struct {

    // errCode
    
    ExCode   int64 `json:"ex_code,omitempty"`
    

    // errMsg
    
    ExMsg   string `json:"ex_msg,omitempty"`
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // POS交易号
    
    OutTradeNo   string `json:"out_trade_no,omitempty"`
    

}
*/

type AlibabaMjOcPayResponse struct {

    // errCode
    ExCode   int64 `json:"ex_code,omitempty"`

    // errMsg
    ExMsg   string `json:"ex_msg,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

    // POS交易号
    OutTradeNo   string `json:"out_trade_no,omitempty"`

}
