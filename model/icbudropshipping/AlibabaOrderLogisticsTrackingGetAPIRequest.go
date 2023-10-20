package icbudropshipping

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaorderlogisticstrackinggetAPIRequest 阿里巴巴订单物流轨迹查询 API请求
// alibaba.order.logistics.tracking.get
//
// 阿里巴巴订单物流轨迹查询
type AlibabaorderlogisticstrackinggetAPIRequest struct {
	model.Params
	// order id
	_tradeId int64
}

// NewAlibabaorderlogisticstrackinggetRequest 初始化AlibabaorderlogisticstrackinggetAPIRequest对象
func NewAlibabaorderlogisticstrackinggetRequest() *AlibabaorderlogisticstrackinggetAPIRequest {
	return &AlibabaorderlogisticstrackinggetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaorderlogisticstrackinggetAPIRequest) GetApiMethodName() string {
	return "alibaba.order.logistics.tracking.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaorderlogisticstrackinggetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaorderlogisticstrackinggetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeId is TradeId Setter
// order id
func (r *AlibabaorderlogisticstrackinggetAPIRequest) SetTradeId(_tradeId int64) error {
	r._tradeId = _tradeId
	r.Set("trade_id", _tradeId)
	return nil
}

// GetTradeId TradeId Getter
func (r AlibabaorderlogisticstrackinggetAPIRequest) GetTradeId() int64 {
	return r._tradeId
}
