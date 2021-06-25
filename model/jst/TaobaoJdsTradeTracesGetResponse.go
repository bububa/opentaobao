package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取单条订单跟踪详情 APIResponse
taobao.jds.trade.traces.get

获取聚石塔数据共享的交易全链路信息
*/
type TaobaoJdsTradeTracesGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoJdsTradeTracesGetResponse `json:"taobao_jds_trade_traces_get_response,omitempty"`
}

type TaobaoJdsTradeTracesGetResponse struct {

    // 跟踪信息列表
    Traces   []TradeTrace `json:"traces,omitempty"`

    // 在订单全链路系统中的状态(比如是否存在)
    UserStatus   string `json:"user_status,omitempty"`

}
