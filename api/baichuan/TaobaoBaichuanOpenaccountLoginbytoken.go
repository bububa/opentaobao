package baichuan

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baichuan"
)

/* 
百川TOKEN 登录 
taobao.baichuan.openaccount.loginbytoken

百川TOKEN 登录
*/
func TaobaoBaichuanOpenaccountLoginbytoken(clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountLoginbytokenRequest, session string) (*baichuan.TaobaoBaichuanOpenaccountLoginbytokenAPIResponse, error) {
    var resp baichuan.TaobaoBaichuanOpenaccountLoginbytokenAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
