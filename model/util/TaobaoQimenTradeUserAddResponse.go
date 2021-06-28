package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
添加奇门订单链路用户 APIResponse
taobao.qimen.trade.user.add

添加奇门订单链路用户
*/
type TaobaoQimenTradeUserAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenTradeUserAddResponse `json:"qimen_trade_user_add_response,omitempty"` 
    TaobaoQimenTradeUserAddResponse
}

/* model for simplify = false
type TaobaoQimenTradeUserAddResponse struct {

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // appkey
    
    Appkey   string `json:"appkey,omitempty"`
    

    // 卖家备注
    
    Memo   string `json:"memo,omitempty"`
    

}
*/

type TaobaoQimenTradeUserAddResponse struct {

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // appkey
    Appkey   string `json:"appkey,omitempty"`

    // 卖家备注
    Memo   string `json:"memo,omitempty"`

}
