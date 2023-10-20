package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamossupplierbasisgetsupplierinfoAPIRequest 获取供应商基础信息 API请求
// alibaba.mos.supplier.basis.getsupplierinfo
//
// 基于供应商id获取供应商基础脱敏信息
type AlibabamossupplierbasisgetsupplierinfoAPIRequest struct {
	model.Params
	// 供应商id
	_supplierId string
}

// NewAlibabamossupplierbasisgetsupplierinfoRequest 初始化AlibabamossupplierbasisgetsupplierinfoAPIRequest对象
func NewAlibabamossupplierbasisgetsupplierinfoRequest() *AlibabamossupplierbasisgetsupplierinfoAPIRequest {
	return &AlibabamossupplierbasisgetsupplierinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamossupplierbasisgetsupplierinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.supplier.basis.getsupplierinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamossupplierbasisgetsupplierinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamossupplierbasisgetsupplierinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSupplierId is SupplierId Setter
// 供应商id
func (r *AlibabamossupplierbasisgetsupplierinfoAPIRequest) SetSupplierId(_supplierId string) error {
	r._supplierId = _supplierId
	r.Set("supplier_id", _supplierId)
	return nil
}

// GetSupplierId SupplierId Getter
func (r AlibabamossupplierbasisgetsupplierinfoAPIRequest) GetSupplierId() string {
	return r._supplierId
}
