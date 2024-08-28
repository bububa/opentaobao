package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemSkuAdd 添加SKU
// taobao.item.sku.add
//
// 新增一个sku到num_iid指定的商品中 &lt;br/&gt;传入的iid所对应的商品必须属于当前会话的用户
func TaobaoItemSkuAdd(ctx context.Context, clt *core.SDKClient, req *tbitem.TaobaoItemSkuAddAPIRequest, resp *tbitem.TaobaoItemSkuAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
