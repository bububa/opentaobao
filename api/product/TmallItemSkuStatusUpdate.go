package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemSkuStatusUpdate 商品sku状态更新
// tmall.item.sku.status.update
//
// 商品sku上下架状态更新
func TmallItemSkuStatusUpdate(ctx context.Context, clt *core.SDKClient, req *product.TmallItemSkuStatusUpdateAPIRequest, resp *product.TmallItemSkuStatusUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
