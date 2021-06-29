package miniappopen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/miniappopen"
)

/* 
（已废弃）构建实例化应用 
taobao.miniappp.template.instantiate

实例化saas化的小程序
*/
func TaobaoMiniapppTemplateInstantiate(clt *core.SDKClient, req *miniappopen.TaobaoMiniapppTemplateInstantiateRequest, session string) (*miniappopen.TaobaoMiniapppTemplateInstantiateAPIResponse, error) {
    var resp miniappopen.TaobaoMiniapppTemplateInstantiateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}