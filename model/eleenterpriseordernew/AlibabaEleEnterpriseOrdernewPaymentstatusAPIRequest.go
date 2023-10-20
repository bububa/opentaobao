package eleenterpriseordernew

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeleenterpriseordernewpaymentstatusAPIRequest 设置订单支付 API请求
// alibaba.ele.enterprise.ordernew.paymentstatus
//
// 设置订单支付成功
type AlibabaeleenterpriseordernewpaymentstatusAPIRequest struct {
	model.Params
	// 订单id
	_orderId string
	// 支付流水号
	_paySerialNumber string
}

// NewAlibabaeleenterpriseordernewpaymentstatusRequest 初始化AlibabaeleenterpriseordernewpaymentstatusAPIRequest对象
func NewAlibabaeleenterpriseordernewpaymentstatusRequest() *AlibabaeleenterpriseordernewpaymentstatusAPIRequest {
	return &AlibabaeleenterpriseordernewpaymentstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeleenterpriseordernewpaymentstatusAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.ordernew.paymentstatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeleenterpriseordernewpaymentstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeleenterpriseordernewpaymentstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabaeleenterpriseordernewpaymentstatusAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaeleenterpriseordernewpaymentstatusAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetPaySerialNumber is PaySerialNumber Setter
// 支付流水号
func (r *AlibabaeleenterpriseordernewpaymentstatusAPIRequest) SetPaySerialNumber(_paySerialNumber string) error {
	r._paySerialNumber = _paySerialNumber
	r.Set("pay_serial_number", _paySerialNumber)
	return nil
}

// GetPaySerialNumber PaySerialNumber Getter
func (r AlibabaeleenterpriseordernewpaymentstatusAPIRequest) GetPaySerialNumber() string {
	return r._paySerialNumber
}
