package icbudropshipping

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbudropshipping"
)

// Alibabadropshippingorderpay alibaba dropshipping 支付代扣
// alibaba.dropshipping.order.pay
//
// alibaba dropshipping 支付代扣
func Alibabadropshippingorderpay(clt *core.SDKClient, req *icbudropshipping.AlibabadropshippingorderpayAPIRequest, session string) (*icbudropshipping.AlibabadropshippingorderpayAPIResponse, error) {
	var resp icbudropshipping.AlibabadropshippingorderpayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
