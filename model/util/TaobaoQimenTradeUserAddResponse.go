package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
添加奇门订单链路用户 APIResponse
taobao.qimen.trade.user.add

添加奇门订单链路用户
*/
type TaobaoQimenTradeUserAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_trade_user_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty" xml:"