package lstlogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstlogisticstracequeryAPIRequest 供应商-异云-查询运单物流追踪信息 API请求
// alibaba.lst.logistics.trace.query
//
// 查询LP单物流追踪信息
type AlibabalstlogisticstracequeryAPIRequest struct {
	model.Params
	// 入参
	_query *LstLogisticsTraceQuery
}

// NewAlibabalstlogisticstracequeryRequest 初始化AlibabalstlogisticstracequeryAPIRequest对象
func NewAlibabalstlogisticstracequeryRequest() *AlibabalstlogisticstracequeryAPIRequest {
	return &AlibabalstlogisticstracequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstlogisticstracequeryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.logistics.trace.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstlogisticstracequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstlogisticstracequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 入参
func (r *AlibabalstlogisticstracequeryAPIRequest) SetQuery(_query *LstLogisticsTraceQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabalstlogisticstracequeryAPIRequest) GetQuery() *LstLogisticsTraceQuery {
	return r._query
}
