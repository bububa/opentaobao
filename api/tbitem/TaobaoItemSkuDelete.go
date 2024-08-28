package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemSkuDelete 删除SKU
// taobao.item.sku.delete
//
// 删除一个sku的数据&lt;br/&gt;需要删除的sku通过属性properties进行匹配查找
func TaobaoItemSkuDelete(ctx context.Context, clt *core.SDKClient, req *tbitem.TaobaoItemSkuDeleteAPIRequest, resp *tbitem.TaobaoItemSkuDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
