package icbudropshipping

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbudropshipping"
)

// AlibabaDropshippingOrderPay alibaba dropshipping 支付代扣
// alibaba.dropshipping.order.pay
//
// alibaba dropshipping 支付代扣
func AlibabaDropshippingOrderPay(clt *core.SDKClient, req *icbudropshipping.AlibabaDropshippingOrderPayAPIRequest, resp *icbudropshipping.AlibabaDropshippingOrderPayAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
