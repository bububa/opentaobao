package drug

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest 阿里健康医保支付信息获取 API请求
// alibaba.alihealth.nr.trade.medical.insurance.get
//
// 阿里健康医保支付信息获取
type AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest struct {
	model.Params
	// 淘宝订单ID
	_orderId int64
}

// NewAlibabaAlihealthNrTradeMedicalInsuranceGetRequest 初始化AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest对象
func NewAlibabaAlihealthNrTradeMedicalInsuranceGetRequest() *AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest {
	return &AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest) Reset() {
	r._orderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.nr.trade.medical.insurance.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 淘宝订单ID
func (r *AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest) GetOrderId() int64 {
	return r._orderId
}

var poolAlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthNrTradeMedicalInsuranceGetRequest()
	},
}

// GetAlibabaAlihealthNrTradeMedicalInsuranceGetRequest 从 sync.Pool 获取 AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest
func GetAlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest() *AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest {
	return poolAlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest.Get().(*AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest)
}

// ReleaseAlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest 将 AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest(v *AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest.Put(v)
}
