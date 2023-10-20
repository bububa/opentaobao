package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkstockrealqueryAPIRequest 仓内实时库存查询 API请求
// alibaba.wdk.stock.real.query
//
// 查询仓内实时库存信息
type AlibabawdkstockrealqueryAPIRequest struct {
	model.Params
	// 系统自动生成
	_query *WmsInventoryTopQuery
}

// NewAlibabawdkstockrealqueryRequest 初始化AlibabawdkstockrealqueryAPIRequest对象
func NewAlibabawdkstockrealqueryRequest() *AlibabawdkstockrealqueryAPIRequest {
	return &AlibabawdkstockrealqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkstockrealqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.stock.real.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkstockrealqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkstockrealqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 系统自动生成
func (r *AlibabawdkstockrealqueryAPIRequest) SetQuery(_query *WmsInventoryTopQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabawdkstockrealqueryAPIRequest) GetQuery() *WmsInventoryTopQuery {
	return r._query
}
