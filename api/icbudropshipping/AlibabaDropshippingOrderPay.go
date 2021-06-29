package icbudropshipping

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbudropshipping"
)

/* 
alibaba dropshipping 支付代扣 
alibaba.dropshipping.order.pay

alibaba dropshipping 支付代扣
*/
func AlibabaDropshippingOrderPay(clt *core.SDKClient, req *icbudropshipping.AlibabaDropshippingOrderPayRequest, session string) (*icbudropshipping.AlibabaDropshippingOrderPayAPIResponse, error) {
    var resp icbudropshipping.AlibabaDropshippingOrderPayAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
