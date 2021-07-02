package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSocialLocaleGetAPIRequest Locale获取接口 API请求
// aliexpress.social.locale.get
//
// 新增Locale获取接口
type AliexpressSocialLocaleGetAPIRequest struct {
	model.Params
}

// NewAliexpressSocialLocaleGetRequest 初始化AliexpressSocialLocaleGetAPIRequest对象
func NewAliexpressSocialLocaleGetRequest() *AliexpressSocialLocaleGetAPIRequest {
	return &AliexpressSocialLocaleGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSocialLocaleGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.social.locale.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSocialLocaleGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
