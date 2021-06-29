package exchange

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询换货订单留言列表 APIResponse
tmall.exchange.messages.get

查询换货订单留言列表
*/
type TmallExchangeMessagesGetAPIResponse struct {
    model.CommonResponse
    TmallExchangeMessagesGetResponse
}

type TmallExchangeMessagesGetResponse struct {
    XMLName xml.Name `xml:"tmall_exchange_messages_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *RefundMessageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
