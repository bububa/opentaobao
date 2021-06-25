package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
积分操作接口 
alibaba.alsc.crm.open.point.operate

同步积分接口
*/
func AlibabaAlscCrmOpenPointOperate(clt *core.SDKClient, req *alsc.AlibabaAlscCrmOpenPointOperateRequest, session string) (*alsc.AlibabaAlscCrmOpenPointOperateResponse, error) {
    var resp alsc.AlibabaAlscCrmOpenPointOperateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
