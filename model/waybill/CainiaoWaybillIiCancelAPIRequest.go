package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaowaybilliicancelAPIRequest 商家取消获取的电子面单号 API请求
// cainiao.waybill.ii.cancel
//
// 面单号有误需要取消的时候，调用该接口取消获取的电子面单。
type CainiaowaybilliicancelAPIRequest struct {
	model.Params
	// 快递公司code
	_cpCode string
	// 电子面单号
	_waybillCode string
}

// NewCainiaowaybilliicancelRequest 初始化CainiaowaybilliicancelAPIRequest对象
func NewCainiaowaybilliicancelRequest() *CainiaowaybilliicancelAPIRequest {
	return &CainiaowaybilliicancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaowaybilliicancelAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.ii.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaowaybilliicancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaowaybilliicancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCpCode is CpCode Setter
// 快递公司code
func (r *CainiaowaybilliicancelAPIRequest) SetCpCode(_cpCode string) error {
	r._cpCode = _cpCode
	r.Set("cp_code", _cpCode)
	return nil
}

// GetCpCode CpCode Getter
func (r CainiaowaybilliicancelAPIRequest) GetCpCode() string {
	return r._cpCode
}

// SetWaybillCode is WaybillCode Setter
// 电子面单号
func (r *CainiaowaybilliicancelAPIRequest) SetWaybillCode(_waybillCode string) error {
	r._waybillCode = _waybillCode
	r.Set("waybill_code", _waybillCode)
	return nil
}

// GetWaybillCode WaybillCode Getter
func (r CainiaowaybilliicancelAPIRequest) GetWaybillCode() string {
	return r._waybillCode
}
