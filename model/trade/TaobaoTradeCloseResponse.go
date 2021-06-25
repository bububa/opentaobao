package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
卖家关闭一笔交易 APIResponse
taobao.trade.close

关闭一笔订单，可以是主订单或子订单。当订单从创建到关闭时间小于10s的时候，会报“CLOSE_TRADE_TOO_FAST”错误。
*/
type TaobaoTradeCloseAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTradeCloseResponse `json:"taobao_trade_close_response,omitempty"`
}

type TaobaoTradeCloseResponse struct {

    // 关闭交易时返回的Trade信息，可用字段有tid和modified
    Trade   *Trade `json:"trade,omitempty"`

}
