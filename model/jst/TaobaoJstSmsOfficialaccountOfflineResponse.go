package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔公众号下线 APIResponse
taobao.jst.sms.officialaccount.offline

聚石塔公众号下线
*/
type TaobaoJstSmsOfficialaccountOfflineAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoJstSmsOfficialaccountOfflineResponse `json:"jst_sms_officialaccount_offline_response,omitempty"` 
    TaobaoJstSmsOfficialaccountOfflineResponse
}

/* model for simplify = false
type TaobaoJstSmsOfficialaccountOfflineResponse struct {

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

type TaobaoJstSmsOfficialaccountOfflineResponse struct {

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
