package aedropshiper

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AE下单API APIResponse
aliexpress.trade.buy.placeorder

A006_INVALID_ACCOUNT_INFO
*/
type AliexpressTradeBuyPlaceorderAPIResponse struct {
    model.CommonResponse
    AliexpressTradeBuyPlaceorderResponse
}

type AliexpressTradeBuyPlaceorderResponse struct {
    XMLName xml.Name `xml:"aliexpress_trade_buy_placeorder_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *PlaceOrderRes4OpenApiDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
