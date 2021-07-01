package icbudropshipping

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDropshippingOrderPayAPIResponse
alibaba dropshipping 支付代扣 API返回值
alibaba.dropshipping.order.pay

alibaba dropshipping 支付代扣 */
type AlibabaDropshippingOrderPayAPIResponse struct {
	model.CommonResponse
	AlibabaDropshippingOrderPayAPIResponseModel
}

// AlibabaDropshippingOrderPayAPIResponseModel is alibaba dropshipping 支付代扣 成功返回结果
type AlibabaDropshippingOrderPayAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dropshipping_order_pay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// response model
	Value *CashierPayResponse `json:"value,omitempty" xml:"value,omitempty"`
}
