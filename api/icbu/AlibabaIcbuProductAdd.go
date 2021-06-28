package icbu

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbu"
)

/* 
发布产品 
alibaba.icbu.product.add

发布商品,支持sourcing/一口价商品，支持英文和多种语言原发商品
*/
func AlibabaIcbuProductAdd(clt *core.SDKClient, req *icbu.AlibabaIcbuProductAddRequest, session string) (*icbu.AlibabaIcbuProductAddAPIResponse, error) {
    var resp icbu.AlibabaIcbuProductAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
