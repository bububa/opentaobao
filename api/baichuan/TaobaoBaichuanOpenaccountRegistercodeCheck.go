package baichuan

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baichuan"
)

/* 
百川检查注册验证码 
taobao.baichuan.openaccount.registercode.check

百川检查注册验证码
*/
func TaobaoBaichuanOpenaccountRegistercodeCheck(clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountRegistercodeCheckRequest, session string) (*baichuan.TaobaoBaichuanOpenaccountRegistercodeCheckAPIResponse, error) {
    var resp baichuan.TaobaoBaichuanOpenaccountRegistercodeCheckAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
