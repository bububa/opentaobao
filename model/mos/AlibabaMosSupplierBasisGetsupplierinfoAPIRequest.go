package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosSupplierBasisGetsupplierinfoAPIRequest 获取供应商基础信息 API请求
// alibaba.mos.supplier.basis.getsupplierinfo
//
// 基于供应商id获取供应商基础脱敏信息
type AlibabaMosSupplierBasisGetsupplierinfoAPIRequest struct {
	model.Params
	// 供应商id
	_supplierId string
}

// NewAlibabaMosSupplierBasisGetsupplierinfoRequest 初始化AlibabaMosSupplierBasisGetsupplierinfoAPIRequest对象
func NewAlibabaMosSupplierBasisGetsupplierinfoRequest() *AlibabaMosSupplierBasisGetsupplierinfoAPIRequest {
	return &AlibabaMosSupplierBasisGetsupplierinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosSupplierBasisGetsupplierinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.supplier.basis.getsupplierinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosSupplierBasisGetsupplierinfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SupplierId Setter
// 供应商id
func (r *AlibabaMosSupplierBasisGetsupplierinfoAPIRequest) SetSupplierId(_supplierId string) error {
	r._supplierId = _supplierId
	r.Set("supplier_id", _supplierId)
	return nil
}

// Get SupplierId Getter
func (r AlibabaMosSupplierBasisGetsupplierinfoAPIRequest) GetSupplierId() string {
	return r._supplierId
}
