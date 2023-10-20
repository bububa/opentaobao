package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenstockqueryAPIRequest 库存查询接口（多条件） API请求
// taobao.qimen.stock.query
//
// ERP调用奇门的接口,查询商品的库存量
type TaobaoqimenstockqueryAPIRequest struct {
	model.Params
	//
	_request *StockQueryRequest
}

// NewTaobaoqimenstockqueryRequest 初始化TaobaoqimenstockqueryAPIRequest对象
func NewTaobaoqimenstockqueryRequest() *TaobaoqimenstockqueryAPIRequest {
	return &TaobaoqimenstockqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenstockqueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.stock.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenstockqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenstockqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenstockqueryAPIRequest) SetRequest(_request *StockQueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenstockqueryAPIRequest) GetRequest() *StockQueryRequest {
	return r._request
}
