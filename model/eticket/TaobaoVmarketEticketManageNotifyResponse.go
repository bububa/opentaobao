package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
主动发起通知接口 APIResponse
taobao.vmarket.eticket.manage.notify

外部合作商家主动发起通知接口
*/
type TaobaoVmarketEticketManageNotifyAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"vmarket_eticket_manage_notify_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 1:成功
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"