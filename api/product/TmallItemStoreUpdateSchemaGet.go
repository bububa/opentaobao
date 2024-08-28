package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemStoreUpdateSchemaGet 天猫门店商品修改规则获取
// tmall.item.store.update.schema.get
//
// 天猫门店商品修改规则获取
func TmallItemStoreUpdateSchemaGet(ctx context.Context, clt *core.SDKClient, req *product.TmallItemStoreUpdateSchemaGetAPIRequest, resp *product.TmallItemStoreUpdateSchemaGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
