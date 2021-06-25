package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取交易确认收货费用 APIResponse
taobao.trade.confirmfee.get

获取交易确认收货费用，可以获取主订单或子订单的确认收货费用
*/
type TaobaoTradeConfirmfeeGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTradeConfirmfeeGetResponse `json:"taobao_trade_confirmfee_get_response,omitempty"`
}

type TaobaoTradeConfirmfeeGetResponse struct {

    // 获取到的交易确认收货费用
    TradeConfirmFee   *TradeConfirmFee `json:"trade_confirm_fee,omitempty"`

}
