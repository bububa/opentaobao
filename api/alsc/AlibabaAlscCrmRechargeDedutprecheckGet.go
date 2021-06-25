package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
储值核销预先校验 
alibaba.alsc.crm.recharge.dedutprecheck.get

储值核销预先校验接口
*/
func AlibabaAlscCrmRechargeDedutprecheckGet(clt *core.SDKClient, req *alsc.AlibabaAlscCrmRechargeDedutprecheckGetRequest, session string) (*alsc.AlibabaAlscCrmRechargeDedutprecheckGetResponse, error) {
    var resp alsc.AlibabaAlscCrmRechargeDedutprecheckGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
