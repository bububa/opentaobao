package user

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/user"
)

/* 
获取用户公钥 
account.aliyuncs.com.GetPubKey.2013-07-01

根据用户的appkey查询用户的pubkey
*/
func AccountAliyuncsComGetPubKey2013_07_01(clt *core.SDKClient, req *user.AccountAliyuncsComGetPubKey2013_07_01Request, session string) (*user.AccountAliyuncsComGetPubKey2013_07_01APIResponse, error) {
    var resp user.AccountAliyuncsComGetPubKey2013_07_01APIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
