package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenitemmappingqueryAPIRequest 前后端商品映射查询接口 API请求
// taobao.qimen.itemmapping.query
//
// 前后端商品映射查询接口
type TaobaoqimenitemmappingqueryAPIRequest struct {
	model.Params
	//
	_request *TaobaoqimenitemmappingqueryRequest
}

// NewTaobaoqimenitemmappingqueryRequest 初始化TaobaoqimenitemmappingqueryAPIRequest对象
func NewTaobaoqimenitemmappingqueryRequest() *TaobaoqimenitemmappingqueryAPIRequest {
	return &TaobaoqimenitemmappingqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenitemmappingqueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.itemmapping.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenitemmappingqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenitemmappingqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenitemmappingqueryAPIRequest) SetRequest(_request *TaobaoqimenitemmappingqueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenitemmappingqueryAPIRequest) GetRequest() *TaobaoqimenitemmappingqueryRequest {
	return r._request
}
