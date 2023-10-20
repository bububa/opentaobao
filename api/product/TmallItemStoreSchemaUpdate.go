package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemStoreSchemaUpdate 天猫门店商品编辑
// tmall.item.store.schema.update
//
// 天猫门店商品编辑
func TmallItemStoreSchemaUpdate(clt *core.SDKClient, req *product.TmallItemStoreSchemaUpdateAPIRequest, resp *product.TmallItemStoreSchemaUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
