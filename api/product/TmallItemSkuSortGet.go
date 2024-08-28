package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemSkuSortGet sku销售属性顺序获取
// tmall.item.sku.sort.get
//
// sku销售属性顺序获取
func TmallItemSkuSortGet(ctx context.Context, clt *core.SDKClient, req *product.TmallItemSkuSortGetAPIRequest, resp *product.TmallItemSkuSortGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
