package icbudropshipping

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDropshippingOrderPayAPIRequest
alibaba dropshipping 支付代扣 API请求
alibaba.dropshipping.order.pay

alibaba dropshipping 支付代扣 */
type AlibabaDropshippingOrderPayAPIRequest struct {
	model.Params
	// request model
	_paramOrderPayRequest *OrderPayRequest
}

// New
