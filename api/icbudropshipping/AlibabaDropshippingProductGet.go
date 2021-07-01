package icbudropshipping

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbudropshipping"
)

/* 
阿里巴巴dropshipping 产品信息获取 
alibaba.dropshipping.product.get

阿里巴巴dropshipping 产品信息获取
*/
func AlibabaDropshippingProductGet(clt *core.SDKClient, req *icbudropshipping.AlibabaDropshippingProductGetAPIRequest, session string) (*icbudropshipping.AlibabaDropshippingProductGetAPIResponse, error) {
    var resp icbudropshipping.AlibabaDropshippingProductGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
