package icbudropshipping

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
alibaba dropshipping 支付代扣 APIResponse
alibaba.dropshipping.order.pay

alibaba dropshipping 支付代扣
*/
type AlibabaDropshippingOrderPayAPIResponse struct {
    model.CommonResponse
    AlibabaDropshippingOrderPayResponse
}

type AlibabaDropshippingOrderPayResponse struct {
    XMLName xml.Name `xml:"alibaba_dropshipping_order_pay_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // response model
    
    Value   *CashierPayResponse `json:"value,omitempty" xml:"value,omitempty"`

    
}
