package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleCarOrderQueryAPIRequest 二手车寄卖查询订单接口 API请求
// alibaba.idle.car.order.query
//
// 二手车寄卖查询订单接口
type AlibabaIdleCarOrderQueryAPIRequest struct {
	model.Params
	// 查询订单入参
	_query *ConsignmentOrderQuery
}

// NewAlibabaIdleCarOrderQueryRequest 初始化AlibabaIdleCarOrderQueryAPIRequest对象
func NewAlibabaIdleCarOrderQueryRequest() *AlibabaIdleCarOrderQueryAPIRequest {
	return &AlibabaIdleCarOrderQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleCarOrderQueryAPIRequest) Reset() {
	r._query = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleCarOrderQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.car.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleCarOrderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleCarOrderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询订单入参
func (r *AlibabaIdleCarOrderQueryAPIRequest) SetQuery(_query *ConsignmentOrderQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabaIdleCarOrderQueryAPIRequest) GetQuery() *ConsignmentOrderQuery {
	return r._query
}

var poolAlibabaIdleCarOrderQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleCarOrderQueryRequest()
	},
}

// GetAlibabaIdleCarOrderQueryRequest 从 sync.Pool 获取 AlibabaIdleCarOrderQueryAPIRequest
func GetAlibabaIdleCarOrderQueryAPIRequest() *AlibabaIdleCarOrderQueryAPIRequest {
	return poolAlibabaIdleCarOrderQueryAPIRequest.Get().(*AlibabaIdleCarOrderQueryAPIRequest)
}

// ReleaseAlibabaIdleCarOrderQueryAPIRequest 将 AlibabaIdleCarOrderQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleCarOrderQueryAPIRequest(v *AlibabaIdleCarOrderQueryAPIRequest) {
	v.Reset()
	poolAlibabaIdleCarOrderQueryAPIRequest.Put(v)
}
