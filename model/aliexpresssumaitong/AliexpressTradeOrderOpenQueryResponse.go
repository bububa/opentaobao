package aliexpresssumaitong

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
Aliexpress开放平台订单查询 APIResponse
aliexpress.trade.order.open.query

Aliexpress开放平台订单信息查询
*/
type AliexpressTradeOrderOpenQueryAPIResponse struct {
    model.CommonResponse
    AliexpressTradeOrderOpenQueryResponse
}

type AliexpressTradeOrderOpenQueryResponse struct {
    XMLName xml.Name `xml:"aliexpress_trade_order_open_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 订单查询接口返回值
    
    Result   *OrderQueryResponse `json:"result,omitempty" xml:"result,omitempty"`

    
}
