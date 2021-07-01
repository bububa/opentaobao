package aesolution

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aesolution"
)

/* 
Get Single Product Info 
aliexpress.solution.product.info.get

Get Single Product Info
*/
func AliexpressSolutionProductInfoGet(clt *core.SDKClient, req *aesolution.AliexpressSolutionProductInfoGetAPIRequest, session string) (*aesolution.AliexpressSolutionProductInfoGetAPIResponse, error) {
    var resp aesolution.AliexpressSolutionProductInfoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
