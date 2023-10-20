package drug

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahealthnrlogisticsdeliverynoupdateAPIRequest 上传订单同城快递单号 API请求
// alibaba.health.nr.logistics.deliveryno.update
//
// 上传订单同城快递单号
type AlibabahealthnrlogisticsdeliverynoupdateAPIRequest struct {
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

// NewAlibabahealthnrlogisticsdeliverynoupdateRequest 初始化AlibabahealthnrlogisticsdeliverynoupdateAPIRequest对象
func NewAlibabahealthnrlogisticsdeliverynoupdateRequest() *AlibabahealthnrlogisticsdeliverynoupdateAPIRequest {
	return &AlibabahealthnrlogisticsdeliverynoupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahealthnrlogisticsdeliverynoupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.health.nr.logistics.deliveryno.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahealthnrlogisticsdeliverynoupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahealthnrlogisticsdeliverynoupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCpCode is CpCode Setter
// 快递公司代码
func (r *AlibabahealthnrlogisticsdeliverynoupdateAPIRequest) SetCpCode(_cpCode string) error {
	r._cpCode = _cpCode
	r.Set("cp_code", _cpCode)
	return nil
}

// GetCpCode CpCode Getter
func (r AlibabahealthnrlogisticsdeliverynoupdateAPIRequest) GetCpCode() string {
	return r._cpCode
}

// SetCourierNo is CourierNo Setter
// 快递单号
func (r *AlibabahealthnrlogisticsdeliverynoupdateAPIRequest) SetCourierNo(_courierNo string) error {
	r._courierNo = _courierNo
	r.Set("courier_no", _courierNo)
	return nil
}

// GetCourierNo CourierNo Getter
func (r AlibabahealthnrlogisticsdeliverynoupdateAPIRequest) GetCourierNo() string {
	return r._courierNo
}

// SetOrderId is OrderId Setter
// 订单ID
func (r *AlibabahealthnrlogisticsdeliverynoupdateAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabahealthnrlogisticsdeliverynoupdateAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetForce is Force Setter
// 是否强制上传，1代表强制，其他值代表需要进行cp_code合法性校验
func (r *AlibabahealthnrlogisticsdeliverynoupdateAPIRequest) SetForce(_force int64) error {
	r._force = _force
	r.Set("force", _force)
	return nil
}

// GetForce Force Getter
func (r AlibabahealthnrlogisticsdeliverynoupdateAPIRequest) GetForce() int64 {
	return r._force
}
