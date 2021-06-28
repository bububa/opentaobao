package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除奇门订单链路用户 APIResponse
taobao.qimen.trade.user.delete

删除奇门订单链路用户
*/
type TaobaoQimenTradeUserDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_trade_user_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // modal
    
    Modal   bool `json:"modal,omitempty" xml:"