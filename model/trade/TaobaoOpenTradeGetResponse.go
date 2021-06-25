package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取单笔交易的部分信息(商家应用使用) APIResponse
taobao.open.trade.get

获取单笔交易的部分信息</br>
1.入参fields中传入buyer_nick ，才能返回buyer_open_id
*/
type TaobaoOpenTradeGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoOpenTradeGetResponse `json:"taobao_open_trade_get_response,omitempty"`
}

type TaobaoOpenTradeGetResponse struct {

    // 搜索到的交易信息列表，返回的Trade和Order中包含的具体信息为入参fields请求的字段信息
    Trade   *Trade `json:"trade,omitempty"`

}
