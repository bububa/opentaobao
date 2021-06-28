package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取供应商基础信息 APIResponse
alibaba.mos.supplier.basis.getsupplierinfo

基于供应商id获取供应商基础脱敏信息
*/
type AlibabaMosSupplierBasisGetsupplierinfoAPIResponse struct {
    model.CommonResponse
    AlibabaMosSupplierBasisGetsupplierinfoResponse
}

type AlibabaMosSupplierBasisGetsupplierinfoResponse struct {
    XMLName xml.Name `xml:"alibaba_mos_supplier_basis_getsupplierinfo_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回数据
    
    Result   *JsonResponse `json:"result,omitempty" xml:"result,omitempty"`

    
}
