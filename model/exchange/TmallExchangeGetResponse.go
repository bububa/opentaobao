package exchange

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单笔换货详情 APIResponse
tmall.exchange.get

获取单笔换货详情
*/
type TmallExchangeGetAPIResponse struct {
    model.CommonResponse
    TmallExchangeGetResponse
}

type TmallExchangeGetResponse struct {
    XMLName xml.Name `xml:"tmall_exchange_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *ExchangeBaseResponse `json:"result,omitempty" xml:"result,omitempty"`

    
}
