package baichuan

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baichuan"
)

/* 
百川用户名密码登录 
taobao.baichuan.openaccount.login

百川用户名密码登录
*/
func TaobaoBaichuanOpenaccountLogin(clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountLoginAPIRequest, session string) (*baichuan.TaobaoBaichuanOpenaccountLoginAPIResponse, error) {
    var resp baichuan.TaobaoBaichuanOpenaccountLoginAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
