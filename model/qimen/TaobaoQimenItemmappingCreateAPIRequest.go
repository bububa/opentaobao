package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenitemmappingcreateAPIRequest 前后端商品映射接口 API请求
// taobao.qimen.itemmapping.create
//
// 前后端商品映射
type TaobaoqimenitemmappingcreateAPIRequest struct {
	model.Params
	//
	_request *RequestDo
}

// NewTaobaoqimenitemmappingcreateRequest 初始化TaobaoqimenitemmappingcreateAPIRequest对象
func NewTaobaoqimenitemmappingcreateRequest() *TaobaoqimenitemmappingcreateAPIRequest {
	return &TaobaoqimenitemmappingcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenitemmappingcreateAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.itemmapping.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenitemmappingcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenitemmappingcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenitemmappingcreateAPIRequest) SetRequest(_request *RequestDo) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenitemmappingcreateAPIRequest) GetRequest() *RequestDo {
	return r._request
}
