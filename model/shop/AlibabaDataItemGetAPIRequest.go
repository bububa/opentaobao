package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadataitemgetAPIRequest 获取商品 API请求
// alibaba.data.item.get
//
// 获取商品信息，作为客户端Weex鉴权的虚拟api
type AlibabadataitemgetAPIRequest struct {
	model.Params
	// 获取商品信息，作为客户端Weex鉴权的虚拟api
	_unNamed string
}

// NewAlibabadataitemgetRequest 初始化AlibabadataitemgetAPIRequest对象
func NewAlibabadataitemgetRequest() *AlibabadataitemgetAPIRequest {
	return &AlibabadataitemgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadataitemgetAPIRequest) GetApiMethodName() string {
	return "alibaba.data.item.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadataitemgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadataitemgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUnNamed is UnNamed Setter
// 获取商品信息，作为客户端Weex鉴权的虚拟api
func (r *AlibabadataitemgetAPIRequest) SetUnNamed(_unNamed string) error {
	r._unNamed = _unNamed
	r.Set("un_named", _unNamed)
	return nil
}

// GetUnNamed UnNamed Getter
func (r AlibabadataitemgetAPIRequest) GetUnNamed() string {
	return r._unNamed
}
