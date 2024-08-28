package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemSkuUpdate 更新SKU信息
// taobao.item.sku.update
//
// *更新一个sku的数据 &lt;br/&gt;*需要更新的sku通过属性properties进行匹配查找 &lt;br/&gt;*商品的数量和价格必须大于等于0 &lt;br/&gt;*sku记录会更新到指定的num_iid对应的商品中 &lt;br/&gt;*num_iid对应的商品必须属于当前的会话用户
func TaobaoItemSkuUpdate(ctx context.Context, clt *core.SDKClient, req *tbitem.TaobaoItemSkuUpdateAPIRequest, resp *tbitem.TaobaoItemSkuUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
