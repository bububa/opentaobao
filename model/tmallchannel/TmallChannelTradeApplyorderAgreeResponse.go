package tmallchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商审核同意采购申请单 APIResponse
tmall.channel.trade.applyorder.agree

供应商审核同意采购申请单
*/
type TmallChannelTradeApplyorderAgreeAPIResponse struct {
    model.CommonResponse
    TmallChannelTradeApplyorderAgreeResponse
}

type TmallChannelTradeApplyorderAgreeResponse struct {
    XMLName xml.Name `xml:"tmall_channel_trade_applyorder_agree_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TmallChannelTradeApplyorderAgreeResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
