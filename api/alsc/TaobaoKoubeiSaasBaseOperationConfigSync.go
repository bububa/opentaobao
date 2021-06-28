package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
商家基础经营设置信息同步 
taobao.koubei.saas.base.operation.config.sync

ISV接入口碑SAAS后, 经营设置数据同步到口碑SAAS
*/
func TaobaoKoubeiSaasBaseOperationConfigSync(clt *core.SDKClient, req *alsc.TaobaoKoubeiSaasBaseOperationConfigSyncRequest, session string) (*alsc.TaobaoKoubeiSaasBaseOperationConfigSyncAPIResponse, error) {
    var resp alsc.TaobaoKoubeiSaasBaseOperationConfigSyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
