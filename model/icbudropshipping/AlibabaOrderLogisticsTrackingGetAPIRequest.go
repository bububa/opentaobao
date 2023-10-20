package icbudropshipping

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOrderLogisticsTrackingGetAPIRequest 阿里巴巴订单物流轨迹查询 API请求
// alibaba.order.logistics.tracking.get
//
// 阿里巴巴订单物流轨迹查询
type AlibabaOrderLogisticsTrackingGetAPIRequest struct {
	model.Params
	// order id
	_tradeId int64
}

// NewAlibabaOrderLogisticsTrackingGetRequest 初始化AlibabaOrderLogisticsTrackingGetAPIRequest对象
func NewAlibabaOrderLogisticsTrackingGetRequest() *AlibabaOrderLogisticsTrackingGetAPIRequest {
	return &AlibabaOrderLogisticsTrackingGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaOrderLogisticsTrackingGetAPIRequest) Reset() {
	r._tradeId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaOrderLogisticsTrackingGetAPIRequest) GetApiMethodName() string {
	return "alibaba.order.logistics.tracking.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaOrderLogisticsTrackingGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaOrderLogisticsTrackingGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeId is TradeId Setter
// order id
func (r *AlibabaOrderLogisticsTrackingGetAPIRequest) SetTradeId(_tradeId int64) error {
	r._tradeId = _tradeId
	r.Set("trade_id", _tradeId)
	return nil
}

// GetTradeId TradeId Getter
func (r AlibabaOrderLogisticsTrackingGetAPIRequest) GetTradeId() int64 {
	return r._tradeId
}

var poolAlibabaOrderLogisticsTrackingGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaOrderLogisticsTrackingGetRequest()
	},
}

// GetAlibabaOrderLogisticsTrackingGetRequest 从 sync.Pool 获取 AlibabaOrderLogisticsTrackingGetAPIRequest
func GetAlibabaOrderLogisticsTrackingGetAPIRequest() *AlibabaOrderLogisticsTrackingGetAPIRequest {
	return poolAlibabaOrderLogisticsTrackingGetAPIRequest.Get().(*AlibabaOrderLogisticsTrackingGetAPIRequest)
}

// ReleaseAlibabaOrderLogisticsTrackingGetAPIRequest 将 AlibabaOrderLogisticsTrackingGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaOrderLogisticsTrackingGetAPIRequest(v *AlibabaOrderLogisticsTrackingGetAPIRequest) {
	v.Reset()
	poolAlibabaOrderLogisticsTrackingGetAPIRequest.Put(v)
}
