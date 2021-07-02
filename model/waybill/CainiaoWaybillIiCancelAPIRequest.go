package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillIiCancelAPIRequest 商家取消获取的电子面单号 API请求
// cainiao.waybill.ii.cancel
//
// 面单号有误需要取消的时候，调用该接口取消获取的电子面单。
type CainiaoWaybillIiCancelAPIRequest struct {
	model.Params
	// 快递公司code
	_cpCode string
	// 电子面单号
	_waybillCode string
}

// NewCainiaoWaybillIiCancelRequest 初始化CainiaoWaybillIiCancelAPIRequest对象
func NewCainiaoWaybillIiCancelRequest() *CainiaoWaybillIiCancelAPIRequest {
	return &CainiaoWaybillIiCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoWaybillIiCancelAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.ii.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoWaybillIiCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CpCode Setter
// 快递公司code
func (r *CainiaoWaybillIiCancelAPIRequest) SetCpCode(_cpCode string) error {
	r._cpCode = _cpCode
	r.Set("cp_code", _cpCode)
	return nil
}

// Get CpCode Getter
func (r CainiaoWaybillIiCancelAPIRequest) GetCpCode() string {
	return r._cpCode
}

// Set is WaybillCode Setter
// 电子面单号
func (r *CainiaoWaybillIiCancelAPIRequest) SetWaybillCode(_waybillCode string) error {
	r._waybillCode = _waybillCode
	r.Set("waybill_code", _waybillCode)
	return nil
}

// Get WaybillCode Getter
func (r CainiaoWaybillIiCancelAPIRequest) GetWaybillCode() string {
	return r._waybillCode
}
