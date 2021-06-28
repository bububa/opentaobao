package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
交易帐务查询 APIResponse
taobao.trade.amount.get

卖家查询该笔交易的资金帐务相关的数据；
1. 只供卖家使用，买家不可使用
2. 可查询所有的状态的交易，但不同状态时交易的相关数据可能会有不同
*/
type TaobaoTradeAmountGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTradeAmountGetResponse `json:"trade_amount_get_response,omitempty"` 
    TaobaoTradeAmountGetResponse
}

/* model for simplify = false
type TaobaoTradeAmountGetResponse struct {

    // 主订单的财务信息详情
    
    TradeAmount  *struct {
        TradeAmount  *TradeAmount `json:"trade_amount,omitempty"`
    } `json:"trade_amount,omitempty"`
    

}
*/

type TaobaoTradeAmountGetResponse struct {

    // 主订单的财务信息详情
    TradeAmount   *TradeAmount `json:"trade_amount,omitempty"`

}
