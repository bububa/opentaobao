package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单条订单跟踪详情 APIResponse
taobao.jds.trade.traces.get

获取聚石塔数据共享的交易全链路信息
*/
type TaobaoJdsTradeTracesGetAPIResponse struct {
    model.CommonResponse
    TaobaoJdsTradeTracesGetResponse
}

type TaobaoJdsTradeTracesGetResponse struct {
    XMLName xml.Name `xml:"jds_trade_traces_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 跟踪信息列表
    
    Traces   []TradeTrace `json:"traces,omitempty" xml:"traces>trade_trace,omitempty"`
    
    
    // 在订单全链路系统中的状态(比如是否存在)
    
    UserStatus   string `json:"user_status,omitempty" xml:"user_status,omitempty"`

    
}
