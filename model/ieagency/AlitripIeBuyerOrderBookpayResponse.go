package ieagency

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票】下单预定支付 APIResponse
alitrip.ie.buyer.order.bookpay

【国际机票】 生单预定支付接口
*/
type AlitripIeBuyerOrderBookpayAPIResponse struct {
    model.CommonResponse
    AlitripIeBuyerOrderBookpayResponse
}

type AlitripIeBuyerOrderBookpayResponse struct {
    XMLName xml.Name `xml:"alitrip_ie_buyer_order_bookpay_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 响应
    
    BookPayOrderResult   *BaseApiResult `json:"book_pay_order_result,omitempty" xml:"book_pay_order_result,omitempty"`

    
}
