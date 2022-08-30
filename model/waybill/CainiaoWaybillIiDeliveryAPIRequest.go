package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillIiDeliveryAPIRequest 派件通知接口 API请求
// cainiao.waybill.ii.delivery
//
// 极效前置场景下的使用此接口，通知进行派件
type CainiaoWaybillIiDeliveryAPIRequest struct {
	model.Params
	// 物流供应商编码
	_cpCode string
	// 面单号
	_waybillCode string
	// 派送类型，1:通知派送； -1: 通知退回
	_deliveryAction int64
}

// NewCainiaoWaybillIiDeliveryRequest 初始化CainiaoWaybillIiDeliveryAPIRequest对象
func NewCainiaoWaybillIiDeliveryRequest() *CainiaoWaybillIiDeliveryAPIRequest {
	return &CainiaoWaybillIiDeliveryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoWaybillIiDeliveryAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.ii.delivery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoWaybillIiDeliveryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCpCode is CpCode Setter
// 物流供应商编码
func (r *CainiaoWaybillIiDeliveryAPIRequest) SetCpCode(_cpCode string) error {
	r._cpCode = _cpCode
	r.Set("cp_code", _cpCode)
	return nil
}

// GetCpCode CpCode Getter
func (r CainiaoWaybillIiDeliveryAPIRequest) GetCpCode() string {
	return r._cpCode
}

// SetWaybillCode is WaybillCode Setter
// 面单号
func (r *CainiaoWaybillIiDeliveryAPIRequest) SetWaybillCode(_waybillCode string) error {
	r._waybillCode = _waybillCode
	r.Set("waybill_code", _waybillCode)
	return nil
}

// GetWaybillCode WaybillCode Getter
func (r CainiaoWaybillIiDeliveryAPIRequest) GetWaybillCode() string {
	return r._waybillCode
}

// SetDeliveryAction is DeliveryAction Setter
// 派送类型，1:通知派送； -1: 通知退回
func (r *CainiaoWaybillIiDeliveryAPIRequest) SetDeliveryAction(_deliveryAction int64) error {
	r._deliveryAction = _deliveryAction
	r.Set("delivery_action", _deliveryAction)
	return nil
}

// GetDeliveryAction DeliveryAction Getter
func (r CainiaoWaybillIiDeliveryAPIRequest) GetDeliveryAction() int64 {
	return r._deliveryAction
}
