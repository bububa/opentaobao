package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
储值操作接口 
alibaba.alsc.crm.open.recharge.operate

储值操作接口
*/
func AlibabaAlscCrmOpenRechargeOperate(clt *core.SDKClient, req *alsc.AlibabaAlscCrmOpenRechargeOperateRequest, session string) (*alsc.AlibabaAlscCrmOpenRechargeOperateResponse, error) {
    var resp alsc.AlibabaAlscCrmOpenRechargeOperateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
