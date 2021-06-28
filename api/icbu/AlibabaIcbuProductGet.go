package icbu

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbu"
)

/* 
获得单个商品详情 
alibaba.icbu.product.get

获取商品详情
*/
func AlibabaIcbuProductGet(clt *core.SDKClient, req *icbu.AlibabaIcbuProductGetRequest, session string) (*icbu.AlibabaIcbuProductGetAPIResponse, error) {
    var resp icbu.AlibabaIcbuProductGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
