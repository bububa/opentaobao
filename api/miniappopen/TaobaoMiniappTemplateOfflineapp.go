package miniappopen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/miniappopen"
)

/* 
下线实例化应用 
taobao.miniapp.template.offlineapp

对指定的实例化小程序进行下线,需要指定clients和app_version
*/
func TaobaoMiniappTemplateOfflineapp(clt *core.SDKClient, req *miniappopen.TaobaoMiniappTemplateOfflineappRequest, session string) (*miniappopen.TaobaoMiniappTemplateOfflineappAPIResponse, error) {
    var resp miniappopen.TaobaoMiniappTemplateOfflineappAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
