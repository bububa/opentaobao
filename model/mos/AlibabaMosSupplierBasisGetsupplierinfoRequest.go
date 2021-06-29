package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取供应商基础信息 API请求
alibaba.mos.supplier.basis.getsupplierinfo

基于供应商id获取供应商基础脱敏信息
*/
type AlibabaMosSupplierBasisGetsupplierinfoRequest struct {
    model.Params
    // 供应商id
    supplierId   string
}

// 初始化AlibabaMosSupplierBasisGetsupplierinfoRequest对象
func NewAlibabaMosSupplierBasisGetsupplierinfoRequest() *AlibabaMosSupplierBasisGetsupplierinfoRequest{
    return &AlibabaMosSupplierBasisGetsupplierinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosSupplierBasisGetsupplierinfoRequest) GetApiMethodName() string {
    return "alibaba.mos.supplier.basis.getsupplierinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosSupplierBasisGetsupplierinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SupplierId Setter
// 供应商id
func (r *AlibabaMosSupplierBasisGetsupplierinfoRequest) SetSupplierId(supplierId string) error {
    r.supplierId = supplierId
    r.Set("supplier_id", supplierId)
    return nil
}

// SupplierId Getter
func (r AlibabaMosSupplierBasisGetsupplierinfoRequest) GetSupplierId() string {
    return r.supplierId
}
