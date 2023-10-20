package lstlogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstlogisticssendinfoqueryAPIRequest 供应商-异云-查询主订单包含的物流单 API请求
// alibaba.lst.logistics.sendinfo.query
//
// 查询主订单包含的物流单
type AlibabalstlogisticssendinfoqueryAPIRequest struct {
	model.Params
	// 入参
	_query *LstLogisticsInfoQuery
}

// NewAlibabalstlogisticssendinfoqueryRequest 初始化AlibabalstlogisticssendinfoqueryAPIRequest对象
func NewAlibabalstlogisticssendinfoqueryRequest() *AlibabalstlogisticssendinfoqueryAPIRequest {
	return &AlibabalstlogisticssendinfoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstlogisticssendinfoqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.logistics.sendinfo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstlogisticssendinfoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstlogisticssendinfoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 入参
func (r *AlibabalstlogisticssendinfoqueryAPIRequest) SetQuery(_query *LstLogisticsInfoQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabalstlogisticssendinfoqueryAPIRequest) GetQuery() *LstLogisticsInfoQuery {
	return r._query
}
