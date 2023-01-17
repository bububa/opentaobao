package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionMerchantProfileGetAPIRequest aliexpress.solution.merchant.profile.get API请求
// aliexpress.solution.merchant.profile.get
//
// API for oversea sellers to obtain the normal information, e.g. store id, registration country code.
type AliexpressSolutionMerchantProfileGetAPIRequest struct {
	model.Params
}

// NewAliexpressSolutionMerchantProfileGetRequest 初始化AliexpressSolutionMerchantProfileGetAPIRequest对象
func NewAliexpressSolutionMerchantProfileGetRequest() *AliexpressSolutionMerchantProfileGetAPIRequest {
	return &AliexpressSolutionMerchantProfileGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionMerchantProfileGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.merchant.profile.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionMerchantProfileGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressSolutionMerchantProfileGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
