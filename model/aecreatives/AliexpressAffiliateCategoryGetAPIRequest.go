package aecreatives

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressaffiliatecategorygetAPIRequest AE流量推广类目信息获取API API请求
// aliexpress.affiliate.category.get
//
// 获取AE流量推广类目的API
type AliexpressaffiliatecategorygetAPIRequest struct {
	model.Params
	// 请求安全签名
	_appSignature string
}

// NewAliexpressaffiliatecategorygetRequest 初始化AliexpressaffiliatecategorygetAPIRequest对象
func NewAliexpressaffiliatecategorygetRequest() *AliexpressaffiliatecategorygetAPIRequest {
	return &AliexpressaffiliatecategorygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressaffiliatecategorygetAPIRequest) GetApiMethodName() string {
	return "aliexpress.affiliate.category.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressaffiliatecategorygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressaffiliatecategorygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppSignature is AppSignature Setter
// 请求安全签名
func (r *AliexpressaffiliatecategorygetAPIRequest) SetAppSignature(_appSignature string) error {
	r._appSignature = _appSignature
	r.Set("app_signature", _appSignature)
	return nil
}

// GetAppSignature AppSignature Getter
func (r AliexpressaffiliatecategorygetAPIRequest) GetAppSignature() string {
	return r._appSignature
}
