package pur

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/pur"
)

/* 
采购供应商订单明细查询接口 
alibaba.ceres.supplier.po.querydetail

采购供应商订单明细查询接口
*/
func AlibabaCeresSupplierPoQuerydetail(clt *core.SDKClient, req *pur.AlibabaCeresSupplierPoQuerydetailRequest, session string) (*pur.AlibabaCeresSupplierPoQuerydetailAPIResponse, error) {
    var resp pur.AlibabaCeresSupplierPoQuerydetailAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
