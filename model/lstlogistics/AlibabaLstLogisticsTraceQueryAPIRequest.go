package lstlogistics

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstLogisticsTraceQueryAPIRequest 供应商-异云-查询运单物流追踪信息 API请求
// alibaba.lst.logistics.trace.query
//
// 查询LP单物流追踪信息
type AlibabaLstLogisticsTraceQueryAPIRequest struct {
	model.Params
	// 入参
	_query *LstLogisticsTraceQuery
}

// NewAlibabaLstLogisticsTraceQueryRequest 初始化AlibabaLstLogisticsTraceQueryAPIRequest对象
func NewAlibabaLstLogisticsTraceQueryRequest() *AlibabaLstLogisticsTraceQueryAPIRequest {
	return &AlibabaLstLogisticsTraceQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstLogisticsTraceQueryAPIRequest) Reset() {
	r._query = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstLogisticsTraceQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.logistics.trace.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstLogisticsTraceQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstLogisticsTraceQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 入参
func (r *AlibabaLstLogisticsTraceQueryAPIRequest) SetQuery(_query *LstLogisticsTraceQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabaLstLogisticsTraceQueryAPIRequest) GetQuery() *LstLogisticsTraceQuery {
	return r._query
}

var poolAlibabaLstLogisticsTraceQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstLogisticsTraceQueryRequest()
	},
}

// GetAlibabaLstLogisticsTraceQueryRequest 从 sync.Pool 获取 AlibabaLstLogisticsTraceQueryAPIRequest
func GetAlibabaLstLogisticsTraceQueryAPIRequest() *AlibabaLstLogisticsTraceQueryAPIRequest {
	return poolAlibabaLstLogisticsTraceQueryAPIRequest.Get().(*AlibabaLstLogisticsTraceQueryAPIRequest)
}

// ReleaseAlibabaLstLogisticsTraceQueryAPIRequest 将 AlibabaLstLogisticsTraceQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstLogisticsTraceQueryAPIRequest(v *AlibabaLstLogisticsTraceQueryAPIRequest) {
	v.Reset()
	poolAlibabaLstLogisticsTraceQueryAPIRequest.Put(v)
}
