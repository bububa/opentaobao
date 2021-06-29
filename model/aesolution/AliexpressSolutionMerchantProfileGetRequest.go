package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.merchant.profile.get APIRequest
aliexpress.solution.merchant.profile.get

API for oversea sellers to obtain the normal information, e.g. store id, registration country code.
*/
type AliexpressSolutionMerchantProfileGetRequest struct {
    model.Params

}

func NewAliexpressSolutionMerchantProfileGetRequest() *AliexpressSolutionMerchantProfileGetRequest{
    return &AliexpressSolutionMerchantProfileGetRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSolutionMerchantProfileGetRequest) GetApiMethodName() string {
    return "aliexpress.solution.merchant.profile.get"
}

func (r AliexpressSolutionMerchantProfileGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


