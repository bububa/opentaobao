package eleenterpriseordernew

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeleenterpriseordernewcancelAPIRequest 订单取消 API请求
// alibaba.ele.enterprise.ordernew.cancel
//
// 订单取消
type AlibabaeleenterpriseordernewcancelAPIRequest struct {
	model.Params
	// 饿了么订单ID
	_orderId string
	// 用户手机号
	_phone string
	// 取消原因(取消时提供)
	_reason string
}

// NewAlibabaeleenterpriseordernewcancelRequest 初始化AlibabaeleenterpriseordernewcancelAPIRequest对象
func NewAlibabaeleenterpriseordernewcancelRequest() *AlibabaeleenterpriseordernewcancelAPIRequest {
	return &AlibabaeleenterpriseordernewcancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeleenterpriseordernewcancelAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.ordernew.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeleenterpriseordernewcancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeleenterpriseordernewcancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 饿了么订单ID
func (r *AlibabaeleenterpriseordernewcancelAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaeleenterpriseordernewcancelAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetPhone is Phone Setter
// 用户手机号
func (r *AlibabaeleenterpriseordernewcancelAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r AlibabaeleenterpriseordernewcancelAPIRequest) GetPhone() string {
	return r._phone
}

// SetReason is Reason Setter
// 取消原因(取消时提供)
func (r *AlibabaeleenterpriseordernewcancelAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r AlibabaeleenterpriseordernewcancelAPIRequest) GetReason() string {
	return r._reason
}
