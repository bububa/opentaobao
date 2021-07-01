package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.merchant.profile.get API请求
aliexpress.solution.merchant.profile.get

API for oversea sellers to obtain the normal information, e.g. store id, registration country code.
*/
type AliexpressSolutionMerchantProfileGetAPIRequest struct {
    model.Params
}

// 初始化AliexpressSolutionMerchantProfileGetAPIRequest对象
func NewAliexpressSolutionMerchantProfileGetRequest() *AliexpressSolutionMerchantProfileGetAPIRequest{
    return &AliexpressSolutionMerchantProfileGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionMerchantProfileGetAPIRequest) GetApiMethodName() string {
    return "aliexpress.solution.merchant.profile.get"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionMerchantProfileGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
