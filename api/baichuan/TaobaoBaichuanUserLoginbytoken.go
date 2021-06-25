package baichuan

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baichuan"
)

/* 
百川手淘信任登录 
taobao.baichuan.user.loginbytoken

百川手淘信任登录
*/
func TaobaoBaichuanUserLoginbytoken(clt *core.SDKClient, req *baichuan.TaobaoBaichuanUserLoginbytokenRequest, session string) (*baichuan.TaobaoBaichuanUserLoginbytokenResponse, error) {
    var resp baichuan.TaobaoBaichuanUserLoginbytokenAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
