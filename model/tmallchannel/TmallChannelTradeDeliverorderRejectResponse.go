package tmallchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商拒绝收货确认单 APIResponse
tmall.channel.trade.deliverorder.reject

供应商拒绝收货确认单
*/
type TmallChannelTradeDeliverorderRejectAPIResponse struct {
    model.CommonResponse
    TmallChannelTradeDeliverorderRejectResponse
}

type TmallChannelTradeDeliverorderRejectResponse struct {
    XMLName xml.Name `xml:"tmall_channel_trade_deliverorder_reject_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TmallChannelTradeDeliverorderRejectResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
