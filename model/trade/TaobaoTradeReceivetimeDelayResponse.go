package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
延长交易收货时间 APIResponse
taobao.trade.receivetime.delay

延长交易收货时间
*/
type TaobaoTradeReceivetimeDelayAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTradeReceivetimeDelayResponse `json:"taobao_trade_receivetime_delay_response,omitempty"`
}

type TaobaoTradeReceivetimeDelayResponse struct {

    // 更新后的交易数据，只包括tid和modified两个字段。
    Trade   *Trade `json:"trade,omitempty"`

}
