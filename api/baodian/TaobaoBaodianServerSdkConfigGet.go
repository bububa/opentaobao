package baodian

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baodian"
)

/* 
获取宝点SDK的配置项（已迁移） 
taobao.baodian.server.sdk.config.get

获取SDK各种配置项（已迁移）
*/
func TaobaoBaodianServerSdkConfigGet(clt *core.SDKClient, req *baodian.TaobaoBaodianServerSdkConfigGetRequest, session string) (*baodian.TaobaoBaodianServerSdkConfigGetAPIResponse, error) {
    var resp baodian.TaobaoBaodianServerSdkConfigGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
