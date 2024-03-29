package category

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressSocialDiscategoryGetAPIRequest) Reset() {
	r._locale = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSocialDiscategoryGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.social.discategory.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSocialDiscategoryGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressSocialDiscategoryGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLocale is Locale Setter
// Locale值，格式为language+&#34;_&#34;+country
func (r *AliexpressSocialDiscategoryGetAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// GetLocale Locale Getter
func (r AliexpressSocialDiscategoryGetAPIRequest) GetLocale() string {
	return r._locale
}

var poolAliexpressSocialDiscategoryGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressSocialDiscategoryGetRequest()
	},
}

// GetAliexpressSocialDiscategoryGetRequest 从 sync.Pool 获取 AliexpressSocialDiscategoryGetAPIRequest
func GetAliexpressSocialDiscategoryGetAPIRequest() *AliexpressSocialDiscategoryGetAPIRequest {
	return poolAliexpressSocialDiscategoryGetAPIRequest.Get().(*AliexpressSocialDiscategoryGetAPIRequest)
}

// ReleaseAliexpressSocialDiscategoryGetAPIRequest 将 AliexpressSocialDiscategoryGetAPIRequest 放入 sync.Pool
func ReleaseAliexpressSocialDiscategoryGetAPIRequest(v *AliexpressSocialDiscategoryGetAPIRequest) {
	v.Reset()
	poolAliexpressSocialDiscategoryGetAPIRequest.Put(v)
}
