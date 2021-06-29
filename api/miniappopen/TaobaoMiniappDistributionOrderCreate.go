package miniappopen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/miniappopen"
)

/* 
创建小程序投放计划 
taobao.miniapp.distribution.order.create

帮助商家，创建小程序的投放计划。
*/
func TaobaoMiniappDistributionOrderCreate(clt *core.SDKClient, req *miniappopen.TaobaoMiniappDistributionOrderCreateRequest, session string) (*miniappopen.TaobaoMiniappDistributionOrderCreateAPIResponse, error) {
    var resp miniappopen.TaobaoMiniappDistributionOrderCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
