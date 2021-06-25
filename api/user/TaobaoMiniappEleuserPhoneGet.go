package user

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/user"
)

/* 
获取饿了么用户信息 
taobao.miniapp.eleuser.phone.get

获取饿了么用户信息
*/
func TaobaoMiniappEleuserPhoneGet(clt *core.SDKClient, req *user.TaobaoMiniappEleuserPhoneGetRequest, session string) (*user.TaobaoMiniappEleuserPhoneGetResponse, error) {
    var resp user.TaobaoMiniappEleuserPhoneGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
