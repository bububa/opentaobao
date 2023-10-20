package icbudropshipping

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOrderPayResultQueryAPIRequest alibaba查询订单支付结果 API请求
// alibaba.order.pay.result.query
//
// alibaba查询订单支付结果
type AlibabaOrderPayResultQueryAPIRequest struct {
	model.Params
	// order id
	_tradeId int64
}

// NewAlibabaOrderPayResultQueryRequest 初始化AlibabaOrderPayResultQueryAPIRequest对象
func NewAlibabaOrderPayResultQueryRequest() *AlibabaOrderPayResultQueryAPIRequest {
	return &AlibabaOrderPayResultQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaOrderPayResultQueryAPIRequest) Reset() {
	r._tradeId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaOrderPayResultQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.order.pay.result.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaOrderPayResultQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaOrderPayResultQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeId is TradeId Setter
// order id
func (r *AlibabaOrderPayResultQueryAPIRequest) SetTradeId(_tradeId int64) error {
	r._tradeId = _tradeId
	r.Set("trade_id", _tradeId)
	return nil
}

// GetTradeId TradeId Getter
func (r AlibabaOrderPayResultQueryAPIRequest) GetTradeId() int64 {
	return r._tradeId
}

var poolAlibabaOrderPayResultQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaOrderPayResultQueryRequest()
	},
}

// GetAlibabaOrderPayResultQueryRequest 从 sync.Pool 获取 AlibabaOrderPayResultQueryAPIRequest
func GetAlibabaOrderPayResultQueryAPIRequest() *AlibabaOrderPayResultQueryAPIRequest {
	return poolAlibabaOrderPayResultQueryAPIRequest.Get().(*AlibabaOrderPayResultQueryAPIRequest)
}

// ReleaseAlibabaOrderPayResultQueryAPIRequest 将 AlibabaOrderPayResultQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaOrderPayResultQueryAPIRequest(v *AlibabaOrderPayResultQueryAPIRequest) {
	v.Reset()
	poolAlibabaOrderPayResultQueryAPIRequest.Put(v)
}
