package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemStoreUpdateSchemaGet 天猫门店商品修改规则获取
// tmall.item.store.update.schema.get
//
// 天猫门店商品修改规则获取
func TmallItemStoreUpdateSchemaGet(clt *core.SDKClient, req *product.TmallItemStoreUpdateSchemaGetAPIRequest, resp *product.TmallItemStoreUpdateSchemaGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
