package aliexpresssumaitong

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresstaxationplatformopengetAPIRequest 平台固定参数获取 API请求
// aliexpress.taxation.platform.open.get
//
// Aliexpress开放平台固定参数获取
type AliexpresstaxationplatformopengetAPIRequest struct {
	model.Params
}

// NewAliexpresstaxationplatformopengetRequest 初始化AliexpresstaxationplatformopengetAPIRequest对象
func NewAliexpresstaxationplatformopengetRequest() *AliexpresstaxationplatformopengetAPIRequest {
	return &AliexpresstaxationplatformopengetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresstaxationplatformopengetAPIRequest) GetApiMethodName() string {
	return "aliexpress.taxation.platform.open.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresstaxationplatformopengetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresstaxationplatformopengetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
