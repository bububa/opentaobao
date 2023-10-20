package aesolution

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressSolutionMerchantProfileGetAPIRequest) Reset() {
	r.Params.ToZero()
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

var poolAliexpressSolutionMerchantProfileGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressSolutionMerchantProfileGetRequest()
	},
}

// GetAliexpressSolutionMerchantProfileGetRequest 从 sync.Pool 获取 AliexpressSolutionMerchantProfileGetAPIRequest
func GetAliexpressSolutionMerchantProfileGetAPIRequest() *AliexpressSolutionMerchantProfileGetAPIRequest {
	return poolAliexpressSolutionMerchantProfileGetAPIRequest.Get().(*AliexpressSolutionMerchantProfileGetAPIRequest)
}

// ReleaseAliexpressSolutionMerchantProfileGetAPIRequest 将 AliexpressSolutionMerchantProfileGetAPIRequest 放入 sync.Pool
func ReleaseAliexpressSolutionMerchantProfileGetAPIRequest(v *AliexpressSolutionMerchantProfileGetAPIRequest) {
	v.Reset()
	poolAliexpressSolutionMerchantProfileGetAPIRequest.Put(v)
}
