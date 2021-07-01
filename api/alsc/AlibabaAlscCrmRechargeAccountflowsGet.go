package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
分页查询储值流水 
alibaba.alsc.crm.recharge.accountflows.get

增加分页查询储值流水接口
*/
func AlibabaAlscCrmRechargeAccountflowsGet(clt *core.SDKClient, req *alsc.AlibabaAlscCrmRechargeAccountflowsGetAPIRequest, session string) (*alsc.AlibabaAlscCrmRechargeAccountflowsGetAPIResponse, error) {
    var resp alsc.AlibabaAlscCrmRechargeAccountflowsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
