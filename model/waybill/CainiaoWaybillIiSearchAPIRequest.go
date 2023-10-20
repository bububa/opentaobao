package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaowaybilliisearchAPIRequest 查询面单服务订购及面单使用情况 API请求
// cainiao.waybill.ii.search
//
// 获取发货地&amp;CP开通状态&amp;账户的使用情况
type CainiaowaybilliisearchAPIRequest struct {
	model.Params
	// 物流公司code
	_cpCode string
}

// NewCainiaowaybilliisearchRequest 初始化CainiaowaybilliisearchAPIRequest对象
func NewCainiaowaybilliisearchRequest() *CainiaowaybilliisearchAPIRequest {
	return &CainiaowaybilliisearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaowaybilliisearchAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.ii.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaowaybilliisearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaowaybilliisearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCpCode is CpCode Setter
// 物流公司code
func (r *CainiaowaybilliisearchAPIRequest) SetCpCode(_cpCode string) error {
	r._cpCode = _cpCode
	r.Set("cp_code", _cpCode)
	return nil
}

// GetCpCode CpCode Getter
func (r CainiaowaybilliisearchAPIRequest) GetCpCode() string {
	return r._cpCode
}
