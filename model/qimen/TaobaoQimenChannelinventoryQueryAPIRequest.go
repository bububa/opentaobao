package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenchannelinventoryqueryAPIRequest 渠道库存查询接口 API请求
// taobao.qimen.channelinventory.query
//
// 渠道库存查询
type TaobaoqimenchannelinventoryqueryAPIRequest struct {
	model.Params
	//
	_request *RequestDo
}

// NewTaobaoqimenchannelinventoryqueryRequest 初始化TaobaoqimenchannelinventoryqueryAPIRequest对象
func NewTaobaoqimenchannelinventoryqueryRequest() *TaobaoqimenchannelinventoryqueryAPIRequest {
	return &TaobaoqimenchannelinventoryqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenchannelinventoryqueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.channelinventory.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenchannelinventoryqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenchannelinventoryqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenchannelinventoryqueryAPIRequest) SetRequest(_request *RequestDo) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenchannelinventoryqueryAPIRequest) GetRequest() *RequestDo {
	return r._request
}
