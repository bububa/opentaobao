package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取供应商基础信息 APIResponse
alibaba.mos.supplier.basis.getsupplierinfo

基于供应商id获取供应商基础脱敏信息
*/
type AlibabaMosSupplierBasisGetsupplierinfoAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMosSupplierBasisGetsupplierinfoResponse `json:"alibaba_mos_supplier_basis_getsupplierinfo_response,omitempty"`
}

type AlibabaMosSupplierBasisGetsupplierinfoResponse struct {

    // 返回数据
    Result   *JsonResponse `json:"result,omitempty"`

}
