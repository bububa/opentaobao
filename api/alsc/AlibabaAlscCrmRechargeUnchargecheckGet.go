package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
储值账户退充值校验 
alibaba.alsc.crm.recharge.unchargecheck.get

储值账户退充值校验接口
*/
func AlibabaAlscCrmRechargeUnchargecheckGet(clt *core.SDKClient, req *alsc.AlibabaAlscCrmRechargeUnchargecheckGetRequest, session string) (*alsc.AlibabaAlscCrmRechargeUnchargecheckGetResponse, error) {
    var resp alsc.AlibabaAlscCrmRechargeUnchargecheckGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
