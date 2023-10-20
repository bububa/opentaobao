package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssolutionmerchantprofilegetAPIRequest aliexpress.solution.merchant.profile.get API请求
// aliexpress.solution.merchant.profile.get
//
// API for oversea sellers to obtain the normal information, e.g. store id, registration country code.
type AliexpresssolutionmerchantprofilegetAPIRequest struct {
	model.Params
}

// NewAliexpresssolutionmerchantprofilegetRequest 初始化AliexpresssolutionmerchantprofilegetAPIRequest对象
func NewAliexpresssolutionmerchantprofilegetRequest() *AliexpresssolutionmerchantprofilegetAPIRequest {
	return &AliexpresssolutionmerchantprofilegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssolutionmerchantprofilegetAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.merchant.profile.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssolutionmerchantprofilegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssolutionmerchantprofilegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
