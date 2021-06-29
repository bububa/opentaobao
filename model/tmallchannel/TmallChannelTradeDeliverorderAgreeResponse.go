package tmallchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商审核通过发货确认 API返回值 
tmall.channel.trade.deliverorder.agree

供应商通过收货确认单
*/
type TmallChannelTradeDeliverorderAgreeAPIResponse struct {
    model.CommonResponse
    TmallChannelTradeDeliverorderAgreeResponse
}

// 供应商审核通过发货确认 成功返回结果
type TmallChannelTradeDeliverorderAgreeResponse struct {
    XMLName xml.Name `xml:"tmall_channel_trade_deliverorder_agree_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TmallChannelTradeDeliverorderAgreeResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
