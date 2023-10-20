package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimensingleitemqueryAPIRequest 商品查询接口 API请求
// taobao.qimen.singleitem.query
//
// 商品查询接口
type TaobaoqimensingleitemqueryAPIRequest struct {
	model.Params
	//
	_request *RequestDo
}

// NewTaobaoqimensingleitemqueryRequest 初始化TaobaoqimensingleitemqueryAPIRequest对象
func NewTaobaoqimensingleitemqueryRequest() *TaobaoqimensingleitemqueryAPIRequest {
	return &TaobaoqimensingleitemqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimensingleitemqueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.singleitem.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimensingleitemqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimensingleitemqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimensingleitemqueryAPIRequest) SetRequest(_request *RequestDo) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimensingleitemqueryAPIRequest) GetRequest() *RequestDo {
	return r._request
}
