package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改交易邮费价格 APIResponse
taobao.trade.postage.update

修改订单邮费接口，通过传入订单编号和邮费价格，修改订单的邮费，返回修改时间modified,邮费post_fee,总费用total_fee。
<br/> <span style="color:red"> API取消加邮费功能通知：http://open.taobao.com/support/announcement_detail.htm?tid=24750</span>
*/
type TaobaoTradePostageUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTradePostageUpdateResponse `json:"trade_postage_update_response,omitempty"` 
    TaobaoTradePostageUpdateResponse
}

/* model for simplify = false
type TaobaoTradePostageUpdateResponse struct {

    // 返回trade类型，其中包含修改时间modified，修改邮费post_fee，修改后的总费用total_fee和买家实付款payment
    
    Trade  *struct {
        Trade  *Trade `json:"trade,omitempty"`
    } `json:"trade,omitempty"`
    

}
*/

type TaobaoTradePostageUpdateResponse struct {

    // 返回trade类型，其中包含修改时间modified，修改邮费post_fee，修改后的总费用total_fee和买家实付款payment
    Trade   *Trade `json:"trade,omitempty"`

}
