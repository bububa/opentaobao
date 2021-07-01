package baichuan

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baichuan"
)

/* 
百川新登录二次验证 
taobao.baichuan.openaccount.newlogindoublecheck

百川新登录二次验证
*/
func TaobaoBaichuanOpenaccountNewlogindoublecheck(clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountNewlogindoublecheckAPIRequest, session string) (*baichuan.TaobaoBaichuanOpenaccountNewlogindoublecheckAPIResponse, error) {
    var resp baichuan.TaobaoBaichuanOpenaccountNewlogindoublecheckAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
