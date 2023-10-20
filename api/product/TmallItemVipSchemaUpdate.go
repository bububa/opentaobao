package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemVipSchemaUpdate 大商家商品编辑接口
// tmall.item.vip.schema.update
//
// 大商家编辑商品
func TmallItemVipSchemaUpdate(clt *core.SDKClient, req *product.TmallItemVipSchemaUpdateAPIRequest, resp *product.TmallItemVipSchemaUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
