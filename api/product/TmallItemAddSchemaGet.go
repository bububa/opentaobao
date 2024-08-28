package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemAddSchemaGet 天猫发布商品规则获取
// tmall.item.add.schema.get
//
// 通过类目以及productId获取商品发布规则；
func TmallItemAddSchemaGet(ctx context.Context, clt *core.SDKClient, req *product.TmallItemAddSchemaGetAPIRequest, resp *product.TmallItemAddSchemaGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
