package drug

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest 上传订单同城快递单号 API请求
// alibaba.health.nr.logistics.deliveryno.update
//
// 上传订单同城快递单号
type AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest struct {
	model.Params
	// 快递公司代码
	_cpCode string
	// 快递单号
	_courierNo string
	// 订单ID
	_orderId int64
	// 是否强制上传，1代表强制，其他值代表需要进行cp_code合法性校验
	_force int64
}

// NewAlibabaHealthNrLogisticsDeliverynoUpdateRequest 初始化AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest对象
func NewAlibabaHealthNrLogisticsDeliverynoUpdateRequest() *AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest {
	return &AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest) Reset() {
	r._cpCode = ""
	r._courierNo = ""
	r._orderId = 0
	r._force = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.health.nr.logistics.deliveryno.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCpCode is CpCode Setter
// 快递公司代码
func (r *AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest) SetCpCode(_cpCode string) error {
	r._cpCode = _cpCode
	r.Set("cp_code", _cpCode)
	return nil
}

// GetCpCode CpCode Getter
func (r AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest) GetCpCode() string {
	return r._cpCode
}

// SetCourierNo is CourierNo Setter
// 快递单号
func (r *AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest) SetCourierNo(_courierNo string) error {
	r._courierNo = _courierNo
	r.Set("courier_no", _courierNo)
	return nil
}

// GetCourierNo CourierNo Getter
func (r AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest) GetCourierNo() string {
	return r._courierNo
}

// SetOrderId is OrderId Setter
// 订单ID
func (r *AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetForce is Force Setter
// 是否强制上传，1代表强制，其他值代表需要进行cp_code合法性校验
func (r *AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest) SetForce(_force int64) error {
	r._force = _force
	r.Set("force", _force)
	return nil
}

// GetForce Force Getter
func (r AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest) GetForce() int64 {
	return r._force
}

var poolAlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHealthNrLogisticsDeliverynoUpdateRequest()
	},
}

// GetAlibabaHealthNrLogisticsDeliverynoUpdateRequest 从 sync.Pool 获取 AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest
func GetAlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest() *AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest {
	return poolAlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest.Get().(*AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest)
}

// ReleaseAlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest 将 AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest(v *AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest) {
	v.Reset()
	poolAlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest.Put(v)
}
