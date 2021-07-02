package category

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSocialDiscategoryGetAPIRequest 展示类目获取接口 API请求
// aliexpress.social.discategory.get
//
// AE展示类目获取接口
type AliexpressSocialDiscategoryGetAPIRequest struct {
	model.Params
	// Locale值，格式为language+"_"+country
	_locale string
}

// NewAliexpressSocialDiscategoryGetRequest 初始化AliexpressSocialDiscategoryGetAPIRequest对象
func NewAliexpressSocialDiscategoryGetRequest() *AliexpressSocialDiscategoryGetAPIRequest {
	return &AliexpressSocialDiscategoryGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSocialDiscategoryGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.social.discategory.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSocialDiscategoryGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Locale Setter
// Locale值，格式为language+"_"+country
func (r *AliexpressSocialDiscategoryGetAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// Get Locale Getter
func (r AliexpressSocialDiscategoryGetAPIRequest) GetLocale() string {
	return r._locale
}
