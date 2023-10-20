package lstlogistics

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstLogisticsSendinfoQueryAPIRequest) Reset() {
	r._query = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstLogisticsSendinfoQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.logistics.sendinfo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstLogisticsSendinfoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstLogisticsSendinfoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaLstLogisticsSendinfoQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstLogisticsSendinfoQueryRequest()
	},
}

// GetAlibabaLstLogisticsSendinfoQueryRequest 从 sync.Pool 获取 AlibabaLstLogisticsSendinfoQueryAPIRequest
func GetAlibabaLstLogisticsSendinfoQueryAPIRequest() *AlibabaLstLogisticsSendinfoQueryAPIRequest {
	return poolAlibabaLstLogisticsSendinfoQueryAPIRequest.Get().(*AlibabaLstLogisticsSendinfoQueryAPIRequest)
}

// ReleaseAlibabaLstLogisticsSendinfoQueryAPIRequest 将 AlibabaLstLogisticsSendinfoQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstLogisticsSendinfoQueryAPIRequest(v *AlibabaLstLogisticsSendinfoQueryAPIRequest) {
	v.Reset()
	poolAlibabaLstLogisticsSendinfoQueryAPIRequest.Put(v)
}
