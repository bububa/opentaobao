package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaowaybilliiproductAPIRequest 商家查询物流商产品类型接口 API请求
// cainiao.waybill.ii.product
//
// 商家可以查询物流商的产品类型和服务能力。
type CainiaowaybilliiproductAPIRequest struct {
	model.Params
	// 快递公司code
	_cpCode string
}

// NewCainiaowaybilliiproductRequest 初始化CainiaowaybilliiproductAPIRequest对象
func NewCainiaowaybilliiproductRequest() *CainiaowaybilliiproductAPIRequest {
	return &CainiaowaybilliiproductAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaowaybilliiproductAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.ii.product"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaowaybilliiproductAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaowaybilliiproductAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCpCode is CpCode Setter
// 快递公司code
func (r *CainiaowaybilliiproductAPIRequest) SetCpCode(_cpCode string) error {
	r._cpCode = _cpCode
	r.Set("cp_code", _cpCode)
	return nil
}

// GetCpCode CpCode Getter
func (r CainiaowaybilliiproductAPIRequest) GetCpCode() string {
	return r._cpCode
}
