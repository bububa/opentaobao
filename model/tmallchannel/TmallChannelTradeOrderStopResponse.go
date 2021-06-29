package tmallchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商停止发货 API返回值 
tmall.channel.trade.order.stop

供应商停止发货
*/
type TmallChannelTradeOrderStopAPIResponse struct {
    model.CommonResponse
    TmallChannelTradeOrderStopResponse
}

// 供应商停止发货 成功返回结果
type TmallChannelTradeOrderStopResponse struct {
    XMLName xml.Name `xml:"tmall_channel_trade_order_stop_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TmallChannelTradeOrderStopResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
