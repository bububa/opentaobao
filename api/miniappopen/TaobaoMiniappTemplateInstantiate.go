package miniappopen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/miniappopen"
)

/* 
构建实例化应用 
taobao.miniapp.template.instantiate

实例化saas化的小程序
*/
func TaobaoMiniappTemplateInstantiate(clt *core.SDKClient, req *miniappopen.TaobaoMiniappTemplateInstantiateRequest, session string) (*miniappopen.TaobaoMiniappTemplateInstantiateAPIResponse, error) {
    var resp miniappopen.TaobaoMiniappTemplateInstantiateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
