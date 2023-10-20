package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenstockoutcreateAPIRequest 出库单创建接口 API请求
// taobao.qimen.stockout.create
//
// taobao.qimen.returnorder.create
type TaobaoqimenstockoutcreateAPIRequest struct {
	model.Params
	//
	_request *StockOutCreateRequest
}

// NewTaobaoqimenstockoutcreateRequest 初始化TaobaoqimenstockoutcreateAPIRequest对象
func NewTaobaoqimenstockoutcreateRequest() *TaobaoqimenstockoutcreateAPIRequest {
	return &TaobaoqimenstockoutcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenstockoutcreateAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.stockout.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenstockoutcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenstockoutcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenstockoutcreateAPIRequest) SetRequest(_request *StockOutCreateRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenstockoutcreateAPIRequest) GetRequest() *StockOutCreateRequest {
	return r._request
}
