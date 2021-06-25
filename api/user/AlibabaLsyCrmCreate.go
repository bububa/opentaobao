package user

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/user"
)

/* 
创建客资 
alibaba.lsy.crm.create

欢客调用该接口进行客资创建
*/
func AlibabaLsyCrmCreate(clt *core.SDKClient, req *user.AlibabaLsyCrmCreateRequest, session string) (*user.AlibabaLsyCrmCreateResponse, error) {
    var resp user.AlibabaLsyCrmCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
