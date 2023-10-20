package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimencombineitemqueryAPIRequest 组合货品关系查询接口 API请求
// taobao.qimen.combineitem.query
//
// 组合货品关系查询
type TaobaoqimencombineitemqueryAPIRequest struct {
	model.Params
	//
	_request *TaobaoqimencombineitemqueryRequest
}

// NewTaobaoqimencombineitemqueryRequest 初始化TaobaoqimencombineitemqueryAPIRequest对象
func NewTaobaoqimencombineitemqueryRequest() *TaobaoqimencombineitemqueryAPIRequest {
	return &TaobaoqimencombineitemqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimencombineitemqueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.combineitem.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimencombineitemqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimencombineitemqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimencombineitemqueryAPIRequest) SetRequest(_request *TaobaoqimencombineitemqueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimencombineitemqueryAPIRequest) GetRequest() *TaobaoqimencombineitemqueryRequest {
	return r._request
}
