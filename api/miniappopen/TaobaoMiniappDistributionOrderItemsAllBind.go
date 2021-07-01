package miniappopen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/miniappopen"
)

/* 
小程序投放-基于投放计划绑定/解绑全店商品 
taobao.miniapp.distribution.order.items.all.bind

提供给使用了投放插件的服务商，可以调用该API实现帮助商家更新已创建的投放单中的绑定商品信息。该接口对应的能力是全店投放，该能力的使用，需要联系平台运营进行人工申请，申请通过以后方可调用成功。
*/
func TaobaoMiniappDistributionOrderItemsAllBind(clt *core.SDKClient, req *miniappopen.TaobaoMiniappDistributionOrderItemsAllBindAPIRequest, session string) (*miniappopen.TaobaoMiniappDistributionOrderItemsAllBindAPIResponse, error) {
    var resp miniappopen.TaobaoMiniappDistributionOrderItemsAllBindAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
