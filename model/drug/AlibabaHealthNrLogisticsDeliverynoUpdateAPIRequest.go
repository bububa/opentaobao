package drug

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest 上传订单同城快递单号 API请求
// alibaba.health.nr.logistics.deliveryno.update
//
// 上传订单同城快递单号
type AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest struct {
	model.Params
	// 订单ID
	_orderId int64
	// 快递公司代码
	_cpCode string
	// 快递单号
	_courierNo string
	// 是否强制上传，1代表强制，其他值代表需要进行cp_code合法性校验
	_force int64
}

// NewAlibabaHealthNrLogisticsDeliverynoUpdateRequest 初始化AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest对象
func NewAlibabaHealthNrLogisticsDeliverynoUpdateRequest() *AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest {
	return &AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.health.nr.logistics.deliveryno.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderId Setter
// 订单ID
func (r *AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// Set is CpCode Setter
// 快递公司代码
func (r *AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest) SetCpCode(_cpCode string) error {
	r._cpCode = _cpCode
	r.Set("cp_code", _cpCode)
	return nil
}

// Get CpCode Getter
func (r AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest) GetCpCode() string {
	return r._cpCode
}

// Set is CourierNo Setter
// 快递单号
func (r *AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest) SetCourierNo(_courierNo string) error {
	r._courierNo = _courierNo
	r.Set("courier_no", _courierNo)
	return nil
}

// Get CourierNo Getter
func (r AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest) GetCourierNo() string {
	return r._courierNo
}

// Set is Force Setter
// 是否强制上传，1代表强制，其他值代表需要进行cp_code合法性校验
func (r *AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest) SetForce(_force int64) error {
	r._force = _force
	r.Set("force", _force)
	return nil
}

// Get Force Getter
func (r AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest) GetForce() int64 {
	return r._force
}
