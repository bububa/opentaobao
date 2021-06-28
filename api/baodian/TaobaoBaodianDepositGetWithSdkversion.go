package baodian

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baodian"
)

/* 
查询用户宝点信息（带sdk版本，已迁移） 
taobao.baodian.deposit.get.with.sdkversion

获取用户宝点信息（带sdk版本，已迁移）
*/
func TaobaoBaodianDepositGetWithSdkversion(clt *core.SDKClient, req *baodian.TaobaoBaodianDepositGetWithSdkversionRequest, session string) (*baodian.TaobaoBaodianDepositGetWithSdkversionAPIResponse, error) {
    var resp baodian.TaobaoBaodianDepositGetWithSdkversionAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
