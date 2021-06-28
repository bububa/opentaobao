package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔取消公众号订购 APIResponse
taobao.jst.sms.officialaccount.cancel

聚石塔取消公众号订购
*/
type TaobaoJstSmsOfficialaccountCancelAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoJstSmsOfficialaccountCancelResponse `json:"jst_sms_officialaccount_cancel_response,omitempty"` 
    TaobaoJstSmsOfficialaccountCancelResponse
}

/* model for simplify = false
type TaobaoJstSmsOfficialaccountCancelResponse struct {

    // 系统异常
    
    ResponseCode   string `json:"response_code,omitempty"`
    

    // 成功
    
    ResponseSuccess   bool `json:"response_success,omitempty"`
    

    // 请求id
    
    ResponseId   string `json:"response_id,omitempty"`
    

    // 成功
    
    Module   bool `json:"module,omitempty"`
    

    // 系统异常
    
    Message   string `json:"message,omitempty"`
    

}
*/

type TaobaoJstSmsOfficialaccountCancelResponse struct {

    // 系统异常
    ResponseCode   string `json:"response_code,omitempty"`

    // 成功
    ResponseSuccess   bool `json:"response_success,omitempty"`

    // 请求id
    ResponseId   string `json:"response_id,omitempty"`

    // 成功
    Module   bool `json:"module,omitempty"`

    // 系统异常
    Message   string `json:"message,omitempty"`

}
