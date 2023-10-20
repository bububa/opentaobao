package eleenterpriseordernew

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest) Reset() {
	r._orderId = ""
	r._paySerialNumber = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.ordernew.paymentstatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetPaySerialNumber is PaySerialNumber Setter
// 支付流水号
func (r *AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest) SetPaySerialNumber(_paySerialNumber string) error {
	r._paySerialNumber = _paySerialNumber
	r.Set("pay_serial_number", _paySerialNumber)
	return nil
}

// GetPaySerialNumber PaySerialNumber Getter
func (r AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest) GetPaySerialNumber() string {
	return r._paySerialNumber
}

var poolAlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEleEnterpriseOrdernewPaymentstatusRequest()
	},
}

// GetAlibabaEleEnterpriseOrdernewPaymentstatusRequest 从 sync.Pool 获取 AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest
func GetAlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest() *AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest {
	return poolAlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest.Get().(*AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest)
}

// ReleaseAlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest 将 AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest 放入 sync.Pool
func ReleaseAlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest(v *AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest) {
	v.Reset()
	poolAlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest.Put(v)
}
