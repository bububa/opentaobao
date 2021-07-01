package pur

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/pur"
)

/* 
po反馈创建 
alibaba.pur.supplier.porespcreate

PO反馈接口
*/
func AlibabaPurSupplierPorespcreate(clt *core.SDKClient, req *pur.AlibabaPurSupplierPorespcreateAPIRequest, session string) (*pur.AlibabaPurSupplierPorespcreateAPIResponse, error) {
    var resp pur.AlibabaPurSupplierPorespcreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
