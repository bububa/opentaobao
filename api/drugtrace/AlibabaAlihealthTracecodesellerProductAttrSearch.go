package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
根据商品id获取商品属性 
alibaba.alihealth.tracecodeseller.product.attr.search

根据商品id获取商品属性
*/
func AlibabaAlihealthTracecodesellerProductAttrSearch(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest, session string) (*drugtrace.AlibabaAlihealthTracecodesellerProductAttrSearchAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthTracecodesellerProductAttrSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
