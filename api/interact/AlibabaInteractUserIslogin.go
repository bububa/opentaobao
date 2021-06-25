package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
校验用户是否已经登录 
alibaba.interact.user.islogin

API的功能是校验用户是否登录，ISV调用接口的时候，通过此接口映射到mtop.interact.user.islogin接口上，因此接口只是做一个给ISV注册调用api的入口，没有发生真实的RPC
*/
func AlibabaInteractUserIslogin(clt *core.SDKClient, req *interact.AlibabaInteractUserIsloginRequest, session string) (*interact.AlibabaInteractUserIsloginResponse, error) {
    var resp interact.AlibabaInteractUserIsloginAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
