package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaowaybilliideliveryAPIRequest 派件通知接口 API请求
// cainiao.waybill.ii.delivery
//
// 极效前置场景下的使用此接口，通知进行派件
type CainiaowaybilliideliveryAPIRequest struct {
	model.Params
	// 物流供应商编码
	_cpCode string
	// 面单号
	_waybillCode string
	// 派送类型，1:通知派送； -1: 通知退回
	_deliveryAction int64
}

// NewCainiaowaybilliideliveryRequest 初始化CainiaowaybilliideliveryAPIRequest对象
func NewCainiaowaybilliideliveryRequest() *CainiaowaybilliideliveryAPIRequest {
	return &CainiaowaybilliideliveryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaowaybilliideliveryAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.ii.delivery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaowaybilliideliveryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaowaybilliideliveryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCpCode is CpCode Setter
// 物流供应商编码
func (r *CainiaowaybilliideliveryAPIRequest) SetCpCode(_cpCode string) error {
	r._cpCode = _cpCode
	r.Set("cp_code", _cpCode)
	return nil
}

// GetCpCode CpCode Getter
func (r CainiaowaybilliideliveryAPIRequest) GetCpCode() string {
	return r._cpCode
}

// SetWaybillCode is WaybillCode Setter
// 面单号
func (r *CainiaowaybilliideliveryAPIRequest) SetWaybillCode(_waybillCode string) error {
	r._waybillCode = _waybillCode
	r.Set("waybill_code", _waybillCode)
	return nil
}

// GetWaybillCode WaybillCode Getter
func (r CainiaowaybilliideliveryAPIRequest) GetWaybillCode() string {
	return r._waybillCode
}

// SetDeliveryAction is DeliveryAction Setter
// 派送类型，1:通知派送； -1: 通知退回
func (r *CainiaowaybilliideliveryAPIRequest) SetDeliveryAction(_deliveryAction int64) error {
	r._deliveryAction = _deliveryAction
	r.Set("delivery_action", _deliveryAction)
	return nil
}

// GetDeliveryAction DeliveryAction Getter
func (r CainiaowaybilliideliveryAPIRequest) GetDeliveryAction() int64 {
	return r._deliveryAction
}
