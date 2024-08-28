package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemSellerGet 获取单个商品详细信息
// taobao.item.seller.get
//
// 获取单个商品的全部信息
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=itemapi&#34; target=&#34;_blank&#34;&gt;点击查看更多商品API说明&lt;/a&gt;&lt;/strong&gt;
func TaobaoItemSellerGet(ctx context.Context, clt *core.SDKClient, req *tbitem.TaobaoItemSellerGetAPIRequest, resp *tbitem.TaobaoItemSellerGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
