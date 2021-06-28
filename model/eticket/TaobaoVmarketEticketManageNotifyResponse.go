package eticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
主动发起通知接口 APIResponse
taobao.vmarket.eticket.manage.notify

外部合作商家主动发起通知接口
*/
type TaobaoVmarketEticketManageNotifyAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoVmarketEticketManageNotifyResponse `json:"vmarket_eticket_manage_notify_response,omitempty"` 
    TaobaoVmarketEticketManageNotifyResponse
}

/* model for simplify = false
type TaobaoVmarketEticketManageNotifyResponse struct {

    // 1:成功
    
    RetCode   int64 `json:"ret_code,omitempty"`
    

}
*/

type TaobaoVmarketEticketManageNotifyResponse struct {

    // 1:成功
    RetCode   int64 `json:"ret_code,omitempty"`

}
