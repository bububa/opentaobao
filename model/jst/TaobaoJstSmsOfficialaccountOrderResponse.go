package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔订购公众号 APIResponse
taobao.jst.sms.officialaccount.order

聚石塔订购公众号接口
*/
type TaobaoJstSmsOfficialaccountOrderAPIResponse struct {
    model.CommonResponse
    Response *TaobaoJstSmsOfficialaccountOrderResponse `json:"taobao_jst_sms_officialaccount_order_response,omitempty"`
}

type TaobaoJstSmsOfficialaccountOrderResponse struct {

    // 系统异常
    ResponseCode   string `json:"response_code,omitempty"`

    // 成功
    ResponseSuccess   bool `json:"response_success,omitempty"`

    // 请求id
    ResponseId   string `json:"response_id,omitempty"`

    // 上报成功
    Module   bool `json:"module,omitempty"`

    // 系统异常
    Message   string `json:"message,omitempty"`

}
