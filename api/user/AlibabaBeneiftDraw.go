package user

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/user"
)

/* 
抽奖接口 
alibaba.beneift.draw

抽奖接口
*/
func AlibabaBeneiftDraw(clt *core.SDKClient, req *user.AlibabaBeneiftDrawRequest, session string) (*user.AlibabaBeneiftDrawResponse, error) {
    var resp user.AlibabaBeneiftDrawAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
