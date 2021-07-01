package icbudropshipping

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbudropshipping"
)

/* 
阿里巴巴买家buynow下单接口 
alibaba.buynow.order.create

阿里巴巴买家下单接口
*/
func AlibabaBuynowOrderCreate(clt *core.SDKClient, req *icbudropshipping.AlibabaBuynowOrderCreateAPIRequest, session string) (*icbudropshipping.AlibabaBuynowOrderCreateAPIResponse, error) {
    var resp icbudropshipping.AlibabaBuynowOrderCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
