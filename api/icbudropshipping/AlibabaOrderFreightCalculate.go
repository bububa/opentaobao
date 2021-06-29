package icbudropshipping

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbudropshipping"
)

/* 
阿里巴巴下单场景运费方案计算 
alibaba.order.freight.calculate

icbu开展 drop shipping 业务，阿里巴巴下单场景运费方案计算
alibaba Create order scenario freight calculation
*/
func AlibabaOrderFreightCalculate(clt *core.SDKClient, req *icbudropshipping.AlibabaOrderFreightCalculateRequest, session string) (*icbudropshipping.AlibabaOrderFreightCalculateAPIResponse, error) {
    var resp icbudropshipping.AlibabaOrderFreightCalculateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
