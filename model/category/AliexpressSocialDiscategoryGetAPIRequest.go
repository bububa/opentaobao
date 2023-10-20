package category

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssocialdiscategorygetAPIRequest 展示类目获取接口 API请求
// aliexpress.social.discategory.get
//
// AE展示类目获取接口
type AliexpresssocialdiscategorygetAPIRequest struct {
	model.Params
	// Locale值，格式为language+"_"+country
	_locale string
}

// NewAliexpresssocialdiscategorygetRequest 初始化AliexpresssocialdiscategorygetAPIRequest对象
func NewAliexpresssocialdiscategorygetRequest() *AliexpresssocialdiscategorygetAPIRequest {
	return &AliexpresssocialdiscategorygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssocialdiscategorygetAPIRequest) GetApiMethodName() string {
	return "aliexpress.social.discategory.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssocialdiscategorygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssocialdiscategorygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLocale is Locale Setter
// Locale值，格式为language+&#34;_&#34;+country
func (r *AliexpresssocialdiscategorygetAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// GetLocale Locale Getter
func (r AliexpresssocialdiscategorygetAPIRequest) GetLocale() string {
	return r._locale
}
