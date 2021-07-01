package aesolution

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aesolution"
)

/* 
aliexpress.solution.merchant.profile.get 
aliexpress.solution.merchant.profile.get

API for oversea sellers to obtain the normal information, e.g. store id, registration country code.
*/
func AliexpressSolutionMerchantProfileGet(clt *core.SDKClient, req *aesolution.AliexpressSolutionMerchantProfileGetAPIRequest, session string) (*aesolution.AliexpressSolutionMerchantProfileGetAPIResponse, error) {
    var resp aesolution.AliexpressSolutionMerchantProfileGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
