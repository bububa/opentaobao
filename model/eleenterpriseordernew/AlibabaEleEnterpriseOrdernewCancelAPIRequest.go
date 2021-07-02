package eleenterpriseordernew

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseOrdernewCancelAPIRequest 订单取消 API请求
// alibaba.ele.enterprise.ordernew.cancel
//
// 订单取消
type AlibabaEleEnterpriseOrdernewCancelAPIRequest struct {
	model.Params
	// 饿了么订单ID
	_orderId string
	// 用户手机号
	_phone string
	// 取消原因(取消时提供)
	_reason string
}

// NewAlibabaEleEnterpriseOrdernewCancelRequest 初始化AlibabaEleEnterpriseOrdernewCancelAPIRequest对象
func NewAlibabaEleEnterpriseOrdernewCancelRequest() *AlibabaEleEnterpriseOrdernewCancelAPIRequest {
	return &AlibabaEleEnterpriseOrdernewCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseOrdernewCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.ordernew.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseOrdernewCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderId is OrderId Setter
// 饿了么订单ID
func (r *AlibabaEleEnterpriseOrdernewCancelAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaEleEnterpriseOrdernewCancelAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetPhone is Phone Setter
// 用户手机号
func (r *AlibabaEleEnterpriseOrdernewCancelAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r AlibabaEleEnterpriseOrdernewCancelAPIRequest) GetPhone() string {
	return r._phone
}

// SetReason is Reason Setter
// 取消原因(取消时提供)
func (r *AlibabaEleEnterpriseOrdernewCancelAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r AlibabaEleEnterpriseOrdernewCancelAPIRequest) GetReason() string {
	return r._reason
}
