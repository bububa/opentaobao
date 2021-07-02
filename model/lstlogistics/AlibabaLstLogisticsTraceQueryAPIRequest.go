package lstlogistics

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstLogisticsTraceQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.logistics.trace.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstLogisticsTraceQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
