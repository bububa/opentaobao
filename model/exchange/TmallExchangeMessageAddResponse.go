package exchange

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家创建换货留言 APIResponse
tmall.exchange.message.add

卖家创建换货留言
*/
type TmallExchangeMessageAddAPIResponse struct {
    model.CommonResponse
    TmallExchangeMessageAddResponse
}

type TmallExchangeMessageAddResponse struct {
    XMLName xml.Name `xml:"tmall_exchange_message_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TmallExchangeMessageAddResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
