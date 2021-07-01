package aesolution

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aesolution"
)

/* 
get order list 
aliexpress.solution.order.get

Get Order List from AliExpress
*/
func AliexpressSolutionOrderGet(clt *core.SDKClient, req *aesolution.AliexpressSolutionOrderGetAPIRequest, session string) (*aesolution.AliexpressSolutionOrderGetAPIResponse, error) {
    var resp aesolution.AliexpressSolutionOrderGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
