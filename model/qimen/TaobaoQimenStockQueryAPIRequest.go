package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenStockQueryAPIRequest 库存查询接口（多条件） API请求
// taobao.qimen.stock.query
//
// ERP调用奇门的接口,查询商品的库存量
type TaobaoQimenStockQueryAPIRequest struct {
	model.Params
	//
	_request *StockQueryRequest
}

// NewTaobaoQimenStockQueryRequest 初始化TaobaoQimenStockQueryAPIRequest对象
func NewTaobaoQimenStockQueryRequest() *TaobaoQimenStockQueryAPIRequest {
	return &TaobaoQimenStockQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenStockQueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.stock.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenStockQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenStockQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenStockQueryAPIRequest) SetRequest(_request *StockQueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenStockQueryAPIRequest) GetRequest() *StockQueryRequest {
	return r._request
}
