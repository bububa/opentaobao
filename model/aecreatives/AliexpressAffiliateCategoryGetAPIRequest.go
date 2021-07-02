package aecreatives

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAffiliateCategoryGetAPIRequest AE流量推广类目信息获取API API请求
// aliexpress.affiliate.category.get
//
// 获取AE流量推广类目的API
type AliexpressAffiliateCategoryGetAPIRequest struct {
	model.Params
	// 请求安全签名
	_appSignature string
}

// NewAliexpressAffiliateCategoryGetRequest 初始化AliexpressAffiliateCategoryGetAPIRequest对象
func NewAliexpressAffiliateCategoryGetRequest() *AliexpressAffiliateCategoryGetAPIRequest {
	return &AliexpressAffiliateCategoryGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAffiliateCategoryGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.affiliate.category.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAffiliateCategoryGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AppSignature Setter
// 请求安全签名
func (r *AliexpressAffiliateCategoryGetAPIRequest) SetAppSignature(_appSignature string) error {
	r._appSignature = _appSignature
	r.Set("app_signature", _appSignature)
	return nil
}

// Get AppSignature Getter
func (r AliexpressAffiliateCategoryGetAPIRequest) GetAppSignature() string {
	return r._appSignature
}
