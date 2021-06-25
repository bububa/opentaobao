package user

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/user"
)

/* 
Open Account数据更新 
taobao.open.account.update

Open Account数据更新
*/
func TaobaoOpenAccountUpdate(clt *core.SDKClient, req *user.TaobaoOpenAccountUpdateRequest, session string) (*user.TaobaoOpenAccountUpdateResponse, error) {
    var resp user.TaobaoOpenAccountUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
