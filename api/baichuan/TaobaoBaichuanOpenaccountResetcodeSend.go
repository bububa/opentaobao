package baichuan

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baichuan"
)

/* 
百川发送找回密码验证码 
taobao.baichuan.openaccount.resetcode.send

百川发送找回密码验证码
*/
func TaobaoBaichuanOpenaccountResetcodeSend(clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountResetcodeSendRequest, session string) (*baichuan.TaobaoBaichuanOpenaccountResetcodeSendAPIResponse, error) {
    var resp baichuan.TaobaoBaichuanOpenaccountResetcodeSendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
