package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
积分兑换校验淘宝账号是否存在 APIResponse
alibaba.wt.order.exchange.partner.checktbuser

积分兑换之前校验淘宝账号是否存在
*/
type AlibabaWtOrderExchangePartnerChecktbuserAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWtOrderExchangePartnerChecktbuserResponse `json:"alibaba_wt_order_exchange_partner_checktbuser_response,omitempty"`
}

type AlibabaWtOrderExchangePartnerChecktbuserResponse struct {

    // 返回值，通过model的值true或者false来判断
    Model   bool `json:"model,omitempty"`

    // 成功
    Desc   string `json:"desc,omitempty"`

    // 成功
    ReturnCode   string `json:"return_code,omitempty"`

    // 接口调用返回成功，真正是否存在号码通过model的返回值来判断
    IsSuccess   bool `json:"is_success,omitempty"`

}
