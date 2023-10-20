package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssociallocalegetAPIRequest Locale获取接口 API请求
// aliexpress.social.locale.get
//
// 新增Locale获取接口
type AliexpresssociallocalegetAPIRequest struct {
	model.Params
}

// NewAliexpresssociallocalegetRequest 初始化AliexpresssociallocalegetAPIRequest对象
func NewAliexpresssociallocalegetRequest() *AliexpresssociallocalegetAPIRequest {
	return &AliexpresssociallocalegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssociallocalegetAPIRequest) GetApiMethodName() string {
	return "aliexpress.social.locale.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssociallocalegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssociallocalegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
