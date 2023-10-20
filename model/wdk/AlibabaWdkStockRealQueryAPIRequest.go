package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkStockRealQueryAPIRequest) Reset() {
	r._query = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkStockRealQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.stock.real.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkStockRealQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkStockRealQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 系统自动生成
func (r *AlibabaWdkStockRealQueryAPIRequest) SetQuery(_query *WmsInventoryTopQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabaWdkStockRealQueryAPIRequest) GetQuery() *WmsInventoryTopQuery {
	return r._query
}

var poolAlibabaWdkStockRealQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkStockRealQueryRequest()
	},
}

// GetAlibabaWdkStockRealQueryRequest 从 sync.Pool 获取 AlibabaWdkStockRealQueryAPIRequest
func GetAlibabaWdkStockRealQueryAPIRequest() *AlibabaWdkStockRealQueryAPIRequest {
	return poolAlibabaWdkStockRealQueryAPIRequest.Get().(*AlibabaWdkStockRealQueryAPIRequest)
}

// ReleaseAlibabaWdkStockRealQueryAPIRequest 将 AlibabaWdkStockRealQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkStockRealQueryAPIRequest(v *AlibabaWdkStockRealQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkStockRealQueryAPIRequest.Put(v)
}
