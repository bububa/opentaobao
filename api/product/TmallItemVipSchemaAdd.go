package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemVipSchemaAdd 大商家商品发布接口
// tmall.item.vip.schema.add
//
// 大商家商品发布接口
func TmallItemVipSchemaAdd(clt *core.SDKClient, req *product.TmallItemVipSchemaAddAPIRequest, resp *product.TmallItemVipSchemaAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
