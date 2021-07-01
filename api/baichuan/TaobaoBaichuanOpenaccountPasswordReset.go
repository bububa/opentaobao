package baichuan

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baichuan"
)

/* 
百川找回密码 
taobao.baichuan.openaccount.password.reset

百川找回密码
*/
func TaobaoBaichuanOpenaccountPasswordReset(clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountPasswordResetAPIRequest, session string) (*baichuan.TaobaoBaichuanOpenaccountPasswordResetAPIResponse, error) {
    var resp baichuan.TaobaoBaichuanOpenaccountPasswordResetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
