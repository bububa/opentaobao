package icbu

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbu"
)

/* 
（新）商品发布新接口 
alibaba.icbu.product.schema.add

提供发布ICBU商品的入口
*/
func AlibabaIcbuProductSchemaAdd(clt *core.SDKClient, req *icbu.AlibabaIcbuProductSchemaAddAPIRequest, session string) (*icbu.AlibabaIcbuProductSchemaAddAPIResponse, error) {
    var resp icbu.AlibabaIcbuProductSchemaAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
