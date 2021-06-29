package aesolution

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aesolution"
)

/* 
get order detail info 
aliexpress.solution.order.info.get

get order detail info
*/
func AliexpressSolutionOrderInfoGet(clt *core.SDKClient, req *aesolution.AliexpressSolutionOrderInfoGetRequest, session string) (*aesolution.AliexpressSolutionOrderInfoGetAPIResponse, error) {
    var resp aesolution.AliexpressSolutionOrderInfoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
