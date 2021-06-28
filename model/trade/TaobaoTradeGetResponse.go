package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单笔交易的部分信息(性能高) APIResponse
taobao.trade.get

获取单笔交易的部分信息
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=tradeapi" target="_blank">点击查看更多交易API说明</a></strong>
*/
type TaobaoTradeGetAPIResponse struct {
    model.CommonResponse
    TaobaoTradeGetResponse
}

type TaobaoTradeGetResponse struct {
    XMLName xml.Name `xml:"trade_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 搜索到的交易信息列表，返回的Trade和Order中包含的具体信息为入参fields请求的字段信息
    
    Trade   *Trade `json:"trade,omitempty" xml:"trade,omitempty"`

    
}
