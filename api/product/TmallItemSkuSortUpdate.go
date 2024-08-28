package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemSkuSortUpdate 商品销售属性排序更新
// tmall.item.sku.sort.update
//
// 商品销售属性排序更新
func TmallItemSkuSortUpdate(ctx context.Context, clt *core.SDKClient, req *product.TmallItemSkuSortUpdateAPIRequest, resp *product.TmallItemSkuSortUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
