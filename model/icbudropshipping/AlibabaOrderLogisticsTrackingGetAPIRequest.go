package icbudropshipping

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaOrderLogisticsTrackingGetAPIRequest) GetApiMethodName() string {
	return "alibaba.order.logistics.tracking.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaOrderLogisticsTrackingGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TradeId Setter
// order id
func (r *AlibabaOrderLogisticsTrackingGetAPIRequest) SetTradeId(_tradeId int64) error {
	r._tradeId = _tradeId
	r.Set("trade_id", _tradeId)
	return nil
}

// Get TradeId Getter
func (r AlibabaOrderLogisticsTrackingGetAPIRequest) GetTradeId() int64 {
	return r._tradeId
}
