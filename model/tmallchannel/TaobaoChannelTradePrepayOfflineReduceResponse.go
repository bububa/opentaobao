package tmallchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道分销供应商上传线下流水预存款（减少） APIResponse
taobao.channel.trade.prepay.offline.reduce

渠道分销供应商上传线下流水预存款（减少）
*/
type TaobaoChannelTradePrepayOfflineReduceAPIResponse struct {
    model.CommonResponse
    TaobaoChannelTradePrepayOfflineReduceResponse
}

type TaobaoChannelTradePrepayOfflineReduceResponse struct {
    XMLName xml.Name `xml:"channel_trade_prepay_offline_reduce_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ResultTopDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
