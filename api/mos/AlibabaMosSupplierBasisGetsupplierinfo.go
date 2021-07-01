package mos

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mos"
)

/* 
获取供应商基础信息 
alibaba.mos.supplier.basis.getsupplierinfo

基于供应商id获取供应商基础脱敏信息
*/
func AlibabaMosSupplierBasisGetsupplierinfo(clt *core.SDKClient, req *mos.AlibabaMosSupplierBasisGetsupplierinfoAPIRequest, session string) (*mos.AlibabaMosSupplierBasisGetsupplierinfoAPIResponse, error) {
    var resp mos.AlibabaMosSupplierBasisGetsupplierinfoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
