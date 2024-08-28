package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemSkuNewGet 查询sku销售属性标新信息
// tmall.item.sku.new.get
//
// 查询sku销售属性标新信息
func TmallItemSkuNewGet(ctx context.Context, clt *core.SDKClient, req *product.TmallItemSkuNewGetAPIRequest, resp *product.TmallItemSkuNewGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
