package aesolution

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aesolution"
)

/* 
Query the sku attribute information belonged to a specific category 
aliexpress.solution.sku.attribute.query

Query the sku attribute information belonged to a specific category, customized for oversea merchants.
*/
func AliexpressSolutionSkuAttributeQuery(clt *core.SDKClient, req *aesolution.AliexpressSolutionSkuAttributeQueryRequest, session string) (*aesolution.AliexpressSolutionSkuAttributeQueryAPIResponse, error) {
    var resp aesolution.AliexpressSolutionSkuAttributeQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
