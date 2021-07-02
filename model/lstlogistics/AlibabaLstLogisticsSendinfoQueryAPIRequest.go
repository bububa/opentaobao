package lstlogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstLogisticsSendinfoQueryAPIRequest 供应商-异云-查询主订单包含的物流单 API请求
// alibaba.lst.logistics.sendinfo.query
//
// 查询主订单包含的物流单
type AlibabaLstLogisticsSendinfoQueryAPIRequest struct {
	model.Params
	// 入参
	_query *LstLogisticsInfoQuery
}

// NewAlibabaLstLogisticsSendinfoQueryRequest 初始化AlibabaLstLogisticsSendinfoQueryAPIRequest对象
func NewAlibabaLstLogisticsSendinfoQueryRequest() *AlibabaLstLogisticsSendinfoQueryAPIRequest {
	return &AlibabaLstLogisticsSendinfoQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstLogisticsSendinfoQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.logistics.sendinfo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstLogisticsSendinfoQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetQuery is Query Setter
// 入参
func (r *AlibabaLstLogisticsSendinfoQueryAPIRequest) SetQuery(_query *LstLogisticsInfoQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabaLstLogisticsSendinfoQueryAPIRequest) GetQuery() *LstLogisticsInfoQuery {
	return r._query
}
