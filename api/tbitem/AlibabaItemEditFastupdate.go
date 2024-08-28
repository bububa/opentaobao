package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// AlibabaItemEditFastupdate 商品编辑增量更新
// alibaba.item.edit.fastupdate
//
// 商品编辑增量更新;
// &lt;br/&gt;该接口编辑sku，只能更新价格、库存等信息，不能新增sku;
// &lt;br/&gt;新增sku用全量接口alibaba.item.edit.submit，先设置销售属性;
func AlibabaItemEditFastupdate(ctx context.Context, clt *core.SDKClient, req *tbitem.AlibabaItemEditFastupdateAPIRequest, resp *tbitem.AlibabaItemEditFastupdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
