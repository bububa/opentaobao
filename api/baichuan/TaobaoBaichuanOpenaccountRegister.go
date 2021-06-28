package baichuan

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baichuan"
)

/* 
百川账号注册 
taobao.baichuan.openaccount.register

百川账号注册
*/
func TaobaoBaichuanOpenaccountRegister(clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountRegisterRequest, session string) (*baichuan.TaobaoBaichuanOpenaccountRegisterAPIResponse, error) {
    var resp baichuan.TaobaoBaichuanOpenaccountRegisterAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
