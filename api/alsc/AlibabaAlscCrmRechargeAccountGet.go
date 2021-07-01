package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
查询储值账户信息 
alibaba.alsc.crm.recharge.account.get

查询储值账户信息接口
*/
func AlibabaAlscCrmRechargeAccountGet(clt *core.SDKClient, req *alsc.AlibabaAlscCrmRechargeAccountGetAPIRequest, session string) (*alsc.AlibabaAlscCrmRechargeAccountGetAPIResponse, error) {
    var resp alsc.AlibabaAlscCrmRechargeAccountGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
