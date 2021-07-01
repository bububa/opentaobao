package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
分销订单查询 
alibaba.retail.commission.order.query

查询商家的分销订单
*/
func AlibabaRetailCommissionOrderQuery(clt *core.SDKClient, req *omniorder.AlibabaRetailCommissionOrderQueryAPIRequest, session string) (*omniorder.AlibabaRetailCommissionOrderQueryAPIResponse, error) {
    var resp omniorder.AlibabaRetailCommissionOrderQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
