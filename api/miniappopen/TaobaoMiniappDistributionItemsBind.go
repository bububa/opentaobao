package miniappopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappDistributionItemsBind 【已废弃】小程序投放-商品绑定/解绑
// taobao.miniapp.distribution.items.bind
//
// 【已废弃，请使用 taobao.miniapp.distribution.order.items.bind 接口】提供给使用了投放插件的服务商，可以调用该API实现帮助商家更新已创建的投放单中的绑定商品信息。
func TaobaoMiniappDistributionItemsBind(ctx context.Context, clt *core.SDKClient, req *miniappopen.TaobaoMiniappDistributionItemsBindAPIRequest, resp *miniappopen.TaobaoMiniappDistributionItemsBindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
