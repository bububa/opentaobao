package miniappopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappDistributionOrderItemsBind 小程序投放-基于投放计划绑定/解绑商品
// taobao.miniapp.distribution.order.items.bind
//
// 提供给使用了投放插件的服务商，可以调用该API实现帮助商家更新已创建的投放单中的绑定商品信息。
func TaobaoMiniappDistributionOrderItemsBind(ctx context.Context, clt *core.SDKClient, req *miniappopen.TaobaoMiniappDistributionOrderItemsBindAPIRequest, resp *miniappopen.TaobaoMiniappDistributionOrderItemsBindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
