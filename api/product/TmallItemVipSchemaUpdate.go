package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemVipSchemaUpdate 大商家商品编辑接口
// tmall.item.vip.schema.update
//
// 大商家编辑商品
func TmallItemVipSchemaUpdate(ctx context.Context, clt *core.SDKClient, req *product.TmallItemVipSchemaUpdateAPIRequest, resp *product.TmallItemVipSchemaUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
