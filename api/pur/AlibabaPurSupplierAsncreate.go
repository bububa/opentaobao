package pur

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/pur"
)

/* 
asn创建 
alibaba.pur.supplier.asncreate

asn创建
*/
func AlibabaPurSupplierAsncreate(clt *core.SDKClient, req *pur.AlibabaPurSupplierAsncreateRequest, session string) (*pur.AlibabaPurSupplierAsncreateAPIResponse, error) {
    var resp pur.AlibabaPurSupplierAsncreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
