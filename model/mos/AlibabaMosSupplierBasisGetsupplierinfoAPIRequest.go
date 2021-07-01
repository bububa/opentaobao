package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosSupplierBasisGetsupplierinfoAPIRequest
获取供应商基础信息 API请求
alibaba.mos.supplier.basis.getsupplierinfo

基于供应商id获取供应商基础脱敏信息 */
type AlibabaMosSupplierBasisGetsupplierinfoAPIRequest struct {
	model.Params
	// 供应商id
	_supplierId string
}

// New
