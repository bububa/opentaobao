package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthShortUrlGetAPIRequest 支付宝短链跳转三方h5通用接口 API请求
// alibaba.alihealth.short.url.get
//
// 支付宝短链跳转三方h5通用接口
type AlibabaAlihealthShortUrlGetAPIRequest struct {
	model.Params
	// 三方h5
	_url string
	// 参数替换列表
	_params []string
}

// NewAlibabaAlihealthShortUrlGetRequest 初始化AlibabaAlihealthShortUrlGetAPIRequest对象
func NewAlibabaAlihealthShortUrlGetRequest() *AlibabaAlihealthShortUrlGetAPIRequest {
	return &AlibabaAlihealthShortUrlGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthShortUrlGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.short.url.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthShortUrlGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUrl is Url Setter
// 三方h5
func (r *AlibabaAlihealthShortUrlGetAPIRequest) SetUrl(_url string) error {
	r._url = _url
	r.Set("url", _url)
	return nil
}

// GetUrl Url Getter
func (r AlibabaAlihealthShortUrlGetAPIRequest) GetUrl() string {
	return r._url
}

// SetParams is Params Setter
// 参数替换列表
func (r *AlibabaAlihealthShortUrlGetAPIRequest) SetParams(_params []string) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// GetParams Params Getter
func (r AlibabaAlihealthShortUrlGetAPIRequest) GetParams() []string {
	return r._params
}
