package aedropshiper

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressTradeDsOrderGetAPIRequest 买家查询订单详情 API请求
// aliexpress.trade.ds.order.get
//
// 买家查询订单详情，用于dropshipper
type AliexpressTradeDsOrderGetAPIRequest struct {
	model.Params
	// 订单查询条件
	_singleOrderQuery *AeopSingleOrderQuery
}

// NewAliexpressTradeDsOrderGetRequest 初始化AliexpressTradeDsOrderGetAPIRequest对象
func NewAliexpressTradeDsOrderGetRequest() *AliexpressTradeDsOrderGetAPIRequest {
	return &AliexpressTradeDsOrderGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressTradeDsOrderGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.trade.ds.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressTradeDsOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressTradeDsOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSingleOrderQuery is SingleOrderQuery Setter
// 订单查询条件
func (r *AliexpressTradeDsOrderGetAPIRequest) SetSingleOrderQuery(_singleOrderQuery *AeopSingleOrderQuery) error {
	r._singleOrderQuery = _singleOrderQuery
	r.Set("single_order_query", _singleOrderQuery)
	return nil
}

// GetSingleOrderQuery SingleOrderQuery Getter
func (r AliexpressTradeDsOrderGetAPIRequest) GetSingleOrderQuery() *AeopSingleOrderQuery {
	return r._singleOrderQuery
}
