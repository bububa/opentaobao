package icbudropshipping

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbudropshipping"
)

/* 
阿里巴巴商品运费计算查询接口 
alibaba.shipping.freight.calculate

阿里巴巴商品运费计算查询接口
*/
func AlibabaShippingFreightCalculate(clt *core.SDKClient, req *icbudropshipping.AlibabaShippingFreightCalculateAPIRequest, session string) (*icbudropshipping.AlibabaShippingFreightCalculateAPIResponse, error) {
    var resp icbudropshipping.AlibabaShippingFreightCalculateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
