package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取供应商基础信息 APIRequest
alibaba.mos.supplier.basis.getsupplierinfo

基于供应商id获取供应商基础脱敏信息
*/
type AlibabaMosSupplierBasisGetsupplierinfoRequest struct {
    model.Params

    // 供应商id
    supplierId   string 

}

func NewAlibabaMosSupplierBasisGetsupplierinfoRequest() *AlibabaMosSupplierBasisGetsupplierinfoRequest{
    return &AlibabaMosSupplierBasisGetsupplierinfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosSupplierBasisGetsupplierinfoRequest) GetApiMethodName() string {
    return "alibaba.mos.supplier.basis.getsupplierinfo"
}

func (r AlibabaMosSupplierBasisGetsupplierinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosSupplierBasisGetsupplierinfoRequest) SetSupplierId(supplierId string) error {
    r.supplierId = supplierId
    r.Set("supplier_id", supplierId)
    return nil
}

func (r AlibabaMosSupplierBasisGetsupplierinfoRequest) GetSupplierId() string {
    return r.supplierId
}

