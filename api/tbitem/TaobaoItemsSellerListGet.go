package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemsSellerListGet 批量获取商品详细信息
// taobao.items.seller.list.get
//
// 批量获取商品详细信息
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=itemapi&#34; target=&#34;_blank&#34;&gt;点击查看更多商品API说明&lt;/a&gt;&lt;/strong&gt;
func TaobaoItemsSellerListGet(ctx context.Context, clt *core.SDKClient, req *tbitem.TaobaoItemsSellerListGetAPIRequest, resp *tbitem.TaobaoItemsSellerListGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
