package aesolution

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aesolution"
)

/* 
aliexpress.solution.batch.product.price.update 
aliexpress.solution.batch.product.price.update

batch product price update operation for oversea sellers
*/
func AliexpressSolutionBatchProductPriceUpdate(clt *core.SDKClient, req *aesolution.AliexpressSolutionBatchProductPriceUpdateRequest, session string) (*aesolution.AliexpressSolutionBatchProductPriceUpdateAPIResponse, error) {
    var resp aesolution.AliexpressSolutionBatchProductPriceUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
