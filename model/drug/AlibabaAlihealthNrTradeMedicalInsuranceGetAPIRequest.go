package drug

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthnrtrademedicalinsurancegetAPIRequest 阿里健康医保支付信息获取 API请求
// alibaba.alihealth.nr.trade.medical.insurance.get
//
// 阿里健康医保支付信息获取
type AlibabaalihealthnrtrademedicalinsurancegetAPIRequest struct {
	model.Params
	// 淘宝订单ID
	_orderId int64
}

// NewAlibabaalihealthnrtrademedicalinsurancegetRequest 初始化AlibabaalihealthnrtrademedicalinsurancegetAPIRequest对象
func NewAlibabaalihealthnrtrademedicalinsurancegetRequest() *AlibabaalihealthnrtrademedicalinsurancegetAPIRequest {
	return &AlibabaalihealthnrtrademedicalinsurancegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthnrtrademedicalinsurancegetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.nr.trade.medical.insurance.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthnrtrademedicalinsurancegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthnrtrademedicalinsurancegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 淘宝订单ID
func (r *AlibabaalihealthnrtrademedicalinsurancegetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaalihealthnrtrademedicalinsurancegetAPIRequest) GetOrderId() int64 {
	return r._orderId
}
