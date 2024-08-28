package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemSkuGet 获取SKU
// taobao.item.sku.get
//
// 获取sku_id所对应的sku数据
// sku_id对应的sku要属于传入的nick对应的卖家
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=itemapi&#34; target=&#34;_blank&#34;&gt;点击查看更多商品API说明&lt;/a&gt;&lt;/strong&gt;
func TaobaoItemSkuGet(ctx context.Context, clt *core.SDKClient, req *tbitem.TaobaoItemSkuGetAPIRequest, resp *tbitem.TaobaoItemSkuGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
