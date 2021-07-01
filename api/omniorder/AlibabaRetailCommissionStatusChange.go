package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
分佣状态变更 
alibaba.retail.commission.status.change

分佣系统，分佣状态变更接口
*/
func AlibabaRetailCommissionStatusChange(clt *core.SDKClient, req *omniorder.AlibabaRetailCommissionStatusChangeAPIRequest, session string) (*omniorder.AlibabaRetailCommissionStatusChangeAPIResponse, error) {
    var resp omniorder.AlibabaRetailCommissionStatusChangeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
