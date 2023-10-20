package lsttrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstTradeOrderQuerychangeAPIRequest 订单id批量查询（品牌商视角） API请求
// alibaba.lst.trade.order.querychange
//
// 根据品牌和时间段查询有变更记录的订单id
type AlibabaLstTradeOrderQuerychangeAPIRequest struct {
	model.Params
	// 入参包装类
	_query *LstOrderQuery
}

// NewAlibabaLstTradeOrderQuerychangeRequest 初始化AlibabaLstTradeOrderQuerychangeAPIRequest对象
func NewAlibabaLstTradeOrderQuerychangeRequest() *AlibabaLstTradeOrderQuerychangeAPIRequest {
	return &AlibabaLstTradeOrderQuerychangeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstTradeOrderQuerychangeAPIRequest) Reset() {
	r._query = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeOrderQuerychangeAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.order.querychange"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeOrderQuerychangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstTradeOrderQuerychangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 入参包装类
func (r *AlibabaLstTradeOrderQuerychangeAPIRequest) SetQuery(_query *LstOrderQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabaLstTradeOrderQuerychangeAPIRequest) GetQuery() *LstOrderQuery {
	return r._query
}

var poolAlibabaLstTradeOrderQuerychangeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstTradeOrderQuerychangeRequest()
	},
}

// GetAlibabaLstTradeOrderQuerychangeRequest 从 sync.Pool 获取 AlibabaLstTradeOrderQuerychangeAPIRequest
func GetAlibabaLstTradeOrderQuerychangeAPIRequest() *AlibabaLstTradeOrderQuerychangeAPIRequest {
	return poolAlibabaLstTradeOrderQuerychangeAPIRequest.Get().(*AlibabaLstTradeOrderQuerychangeAPIRequest)
}

// ReleaseAlibabaLstTradeOrderQuerychangeAPIRequest 将 AlibabaLstTradeOrderQuerychangeAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstTradeOrderQuerychangeAPIRequest(v *AlibabaLstTradeOrderQuerychangeAPIRequest) {
	v.Reset()
	poolAlibabaLstTradeOrderQuerychangeAPIRequest.Put(v)
}
