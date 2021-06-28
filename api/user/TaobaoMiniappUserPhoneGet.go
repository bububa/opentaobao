package user

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/user"
)

/* 
获取当前授权用户手机号码 
taobao.miniapp.user.phone.get

在商家应用中，获取当前授权用户手机号码
*/
func TaobaoMiniappUserPhoneGet(clt *core.SDKClient, req *user.TaobaoMiniappUserPhoneGetRequest, session string) (*user.TaobaoMiniappUserPhoneGetAPIResponse, error) {
    var resp user.TaobaoMiniappUserPhoneGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
