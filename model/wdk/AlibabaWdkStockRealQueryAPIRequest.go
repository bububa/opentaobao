package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkStockRealQueryAPIRequest 仓内实时库存查询 API请求
// alibaba.wdk.stock.real.query
//
// 查询仓内实时库存信息
type AlibabaWdkStockRealQueryAPIRequest struct {
	model.Params
	// 系统自动生成
	_query *WmsInventoryTopQuery
}

// NewAlibabaWdkStockRealQueryRequest 初始化AlibabaWdkStockRealQueryAPIRequest对象
func NewAlibabaWdkStockRealQueryRequest() *AlibabaWdkStockRealQueryAPIRequest {
	return &AlibabaWdkStockRealQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkStockRealQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.stock.real.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkStockRealQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Query Setter
// 系统自动生成
func (r *AlibabaWdkStockRealQueryAPIRequest) SetQuery(_query *WmsInventoryTopQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// Get Query Getter
func (r AlibabaWdkStockRealQueryAPIRequest) GetQuery() *WmsInventoryTopQuery {
	return r._query
}
