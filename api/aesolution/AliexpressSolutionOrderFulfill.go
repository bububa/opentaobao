package aesolution

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aesolution"
)

/* 
fulfill order 
aliexpress.solution.order.fulfill

fulfill order for seller
*/
func AliexpressSolutionOrderFulfill(clt *core.SDKClient, req *aesolution.AliexpressSolutionOrderFulfillAPIRequest, session string) (*aesolution.AliexpressSolutionOrderFulfillAPIResponse, error) {
    var resp aesolution.AliexpressSolutionOrderFulfillAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
