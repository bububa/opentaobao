package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
创建及绑定互动实例 
alibaba.interact.isvadmin.bind

创建互动实例，并绑定奖池
*/
func AlibabaInteractIsvadminBind(clt *core.SDKClient, req *interact.AlibabaInteractIsvadminBindRequest, session string) (*interact.AlibabaInteractIsvadminBindResponse, error) {
    var resp interact.AlibabaInteractIsvadminBindAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
