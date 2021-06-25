package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔公众号上线 APIResponse
taobao.jst.sms.officialaccount.online

聚石塔公众号上线
*/
type TaobaoJstSmsOfficialaccountOnlineAPIResponse struct {
    model.CommonResponse
    Response *TaobaoJstSmsOfficialaccountOnlineResponse `json:"taobao_jst_sms_officialaccount_online_response,omitempty"`
}

type TaobaoJstSmsOfficialaccountOnlineResponse struct {

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
