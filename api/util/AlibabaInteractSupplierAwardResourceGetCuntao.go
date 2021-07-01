package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
权益池信息查询 
alibaba.interact.supplier.award.resource.get.cuntao

农村淘宝营销互动权益池信息查询
*/
func AlibabaInteractSupplierAwardResourceGetCuntao(clt *core.SDKClient, req *util.AlibabaInteractSupplierAwardResourceGetCuntaoAPIRequest, session string) (*util.AlibabaInteractSupplierAwardResourceGetCuntaoAPIResponse, error) {
    var resp util.AlibabaInteractSupplierAwardResourceGetCuntaoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
