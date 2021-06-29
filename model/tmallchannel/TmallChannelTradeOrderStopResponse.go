package tmallchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商停止发货 APIResponse
tmall.channel.trade.order.stop

供应商停止发货
*/
type TmallChannelTradeOrderStopAPIResponse struct {
    model.CommonResponse
    TmallChannelTradeOrderStopResponse
}

type TmallChannelTradeOrderStopResponse struct {
    XMLName xml.Name `xml:"tmall_channel_trade_order_stop_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TmallChannelTradeOrderStopResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
