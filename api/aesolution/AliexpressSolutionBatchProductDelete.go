package aesolution

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aesolution"
)

/* 
aliexpress.solution.batch.product.delete 
aliexpress.solution.batch.product.delete

Product delete API. Please note that there is no reverse way to retrieve the products which have been deleted. Use this API in cautious.
*/
func AliexpressSolutionBatchProductDelete(clt *core.SDKClient, req *aesolution.AliexpressSolutionBatchProductDeleteRequest, session string) (*aesolution.AliexpressSolutionBatchProductDeleteAPIResponse, error) {
    var resp aesolution.AliexpressSolutionBatchProductDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
