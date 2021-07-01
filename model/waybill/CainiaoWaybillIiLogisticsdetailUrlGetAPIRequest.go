package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoWaybillIiLogisticsdetailUrlGetAPIRequest
电子面单物流详情授权url获取 API请求
cainiao.waybill.ii.logisticsdetail.url.get

获取电子面单物流详情授权访问的H5 url */
type CainiaoWaybillIiLogisticsdetailUrlGetAPIRequest struct {
	model.Params
	// 快递公司编码
	_cpCode string
	// 电子面单单号
	_waybillCode string
}

// NewCainiaoWaybillIiLogisticsdetailUrlGetRequest 初始化CainiaoWaybillIiLogisticsdetailUrlGetAPIRequest对象
func NewCainiaoWaybillIiLogisticsdetailUrlGetRequest() *CainiaoWaybillIiLogisticsdetailUrlGetAPIRequest {
	return &CainiaoWaybillIiLogisticsdetailUrlGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoWaybillIiLogisticsdetailUrlGetAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.ii.logisticsdetail.url.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoWaybillIiLogisticsdetailUrlGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CpCode Setter
// 快递公司编码
func (r *CainiaoWaybillIiLogisticsdetailUrlGetAPIRequest) SetCpCode(_cpCode string) error {
	r._cpCode = _cpCode
	r.Set("cp_code", _cpCode)
	return nil
}

// Get CpCode Getter
func (r CainiaoWaybillIiLogisticsdetailUrlGetAPIRequest) GetCpCode() string {
	return r._cpCode
}

// Set is WaybillCode Setter
// 电子面单单号
func (r *CainiaoWaybillIiLogisticsdetailUrlGetAPIRequest) SetWaybillCode(_waybillCode string) error {
	r._waybillCode = _waybillCode
	r.Set("waybill_code", _waybillCode)
	return nil
}

// Get WaybillCode Getter
func (r CainiaoWaybillIiLogisticsdetailUrlGetAPIRequest) GetWaybillCode() string {
	return r._waybillCode
}
