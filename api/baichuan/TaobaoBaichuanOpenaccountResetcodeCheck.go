package baichuan

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baichuan"
)

/* 
百川验证找回密码验证码 
taobao.baichuan.openaccount.resetcode.check

百川验证找回密码验证码
*/
func TaobaoBaichuanOpenaccountResetcodeCheck(clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountResetcodeCheckRequest, session string) (*baichuan.TaobaoBaichuanOpenaccountResetcodeCheckAPIResponse, error) {
    var resp baichuan.TaobaoBaichuanOpenaccountResetcodeCheckAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
