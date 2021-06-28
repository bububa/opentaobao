package baichuan

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baichuan"
)

/* 
百川H5登录 
taobao.baichuan.user.login

百川H5登录
*/
func TaobaoBaichuanUserLogin(clt *core.SDKClient, req *baichuan.TaobaoBaichuanUserLoginRequest, session string) (*baichuan.TaobaoBaichuanUserLoginAPIResponse, error) {
    var resp baichuan.TaobaoBaichuanUserLoginAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
