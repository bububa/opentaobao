package eleenterpriseordernew

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest 设置订单支付 API请求
// alibaba.ele.enterprise.ordernew.paymentstatus
//
// 设置订单支付成功
type AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest struct {
	model.Params
	// 订单id
	_orderId string
	// 支付流水号
	_paySerialNumber string
}

// NewAlibabaEleEnterpriseOrdernewPaymentstatusRequest 初始化AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest对象
func NewAlibabaEleEnterpriseOrdernewPaymentstatusRequest() *AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest {
	return &AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.ordernew.paymentstatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderId Setter
// 订单id
func (r *AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest) GetOrderId() string {
	return r._orderId
}

// Set is PaySerialNumber Setter
// 支付流水号
func (r *AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest) SetPaySerialNumber(_paySerialNumber string) error {
	r._paySerialNumber = _paySerialNumber
	r.Set("pay_serial_number", _paySerialNumber)
	return nil
}

// Get PaySerialNumber Getter
func (r AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest) GetPaySerialNumber() string {
	return r._paySerialNumber
}
