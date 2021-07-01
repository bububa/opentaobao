package aliexpresssumaitong

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
Aliexpress开放平台下单前置检查 API返回值 
aliexpress.trade.order.open.check

Aliexpress开放平台下单前通过下单入参获取token
*/
type AliexpressTradeOrderOpenCheckAPIResponse struct {
    model.CommonResponse
    AliexpressTradeOrderOpenCheckAPIResponseModel
}

// Aliexpress开放平台下单前置检查 成功返回结果
type AliexpressTradeOrderOpenCheckAPIResponseModel struct {
    XMLName xml.Name `xml:"aliexpress_trade_order_open_check_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 预下单返回值
    Result   *PreCheckResponse `json:"result,omitempty" xml:"result,omitempty"`
}
